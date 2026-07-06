import { useState, useEffect } from 'react';
import { apiListInmates, apiCreateInmate, apiUpdateInmate, apiDeleteInmate, apiGetInmate, apiListUFs, apiListPrisonUnits } from '../api';
import type { InmatesList, UF, PrisonUnit } from '../types';
import Layout from '../components/Layout';
import Button from '../components/Button';
import Modal from '../components/Modal';
import Input from '../components/Input';
import Select from '../components/Select';
import PhoneInput from '../components/PhoneInput';

export default function Inmates() {
  const [inmates, setInmates] = useState<InmatesList[]>([]);
  const [ufs, setUfs] = useState<UF[]>([]);
  const [originUnits, setOriginUnits] = useState<PrisonUnit[]>([]);
  const [destinationUnits, setDestinationUnits] = useState<PrisonUnit[]>([]);
  const [loading, setLoading] = useState(true);
  const [modalOpen, setModalOpen] = useState(false);
  const [editingId, setEditingId] = useState<number | null>(null);
  const [formData, setFormData] = useState({
    origin_unit_id: 0,
    custody: 'CLOSED' as 'CLOSED' | 'SEMI_OPEN' | 'OPEN',
    destination_unit_ids: [] as number[],
    responsible: { attorney: '', phone: '' }
  });
  const [selectedDestinations, setSelectedDestinations] = useState<PrisonUnit[]>([]);
  const [selectedOriginUF, setSelectedOriginUF] = useState('');
  const [selectedDestinationUF, setSelectedDestinationUF] = useState('');
  const [pendingDestinationUnitId, setPendingDestinationUnitId] = useState(0);
  const [error, setError] = useState('');

  useEffect(() => {
    loadInmates();
    loadUFs();
  }, []);

  const loadUFs = async () => {
    try {
      const data = await apiListUFs();
      setUfs(data ?? []);
    } catch (err: any) {
      console.error('Erro ao carregar UFs:', err);
    }
  };

  const loadOriginUnits = async (ufCode: string) => {
    try {
      const data = await apiListPrisonUnits({ uf_code: ufCode });
      setOriginUnits(data ?? []);
    } catch (err: any) {
      console.error('Erro ao carregar unidades de origem:', err);
    }
  };

  const loadDestinationUnits = async (ufCode: string) => {
    try {
      const data = await apiListPrisonUnits({ uf_code: ufCode });
      setDestinationUnits(data ?? []);
    } catch (err: any) {
      console.error('Erro ao carregar unidades de destino:', err);
    }
  };

  const loadInmates = async () => {
    try {
      setLoading(true);
      const data = await apiListInmates();
      setInmates(data);
    } catch (err: any) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError('');

    if (formData.destination_unit_ids.length === 0) {
      setError('Adicione pelo menos uma unidade prisional de destino.');
      return;
    }

    try {
      if (editingId) {
        await apiUpdateInmate(editingId, formData);
      } else {
        await apiCreateInmate(formData);
      }
      setModalOpen(false);
      resetForm();
      loadInmates();
    } catch (err: any) {
      setError(err.message);
    }
  };

  const handleEdit = async (id: number) => {
    try {
      const inmate = await apiGetInmate(id);
      const destinations = inmate.destination_units ?? [];

      setEditingId(id);
      setFormData({
        origin_unit_id: inmate.origin_unit_id,
        custody: inmate.custody,
        destination_unit_ids: inmate.destination_unit_ids ?? destinations.map((unit) => unit.id),
        responsible: inmate.responsible
      });
      setSelectedDestinations(destinations);

      if (inmate.origin_unit?.uf_code) {
        setSelectedOriginUF(inmate.origin_unit.uf_code);
        await loadOriginUnits(inmate.origin_unit.uf_code);
      }

      setModalOpen(true);
    } catch (err: any) {
      setError(err.message);
    }
  };

  const handleDelete = async (id: number) => {
    if (!confirm('Tem certeza que deseja excluir este reeducando?')) return;
    try {
      await apiDeleteInmate(id);
      loadInmates();
    } catch (err: any) {
      setError(err.message);
    }
  };

  const resetForm = () => {
    setEditingId(null);
    setFormData({
      origin_unit_id: 0,
      custody: 'CLOSED',
      destination_unit_ids: [],
      responsible: { attorney: '', phone: '' }
    });
    setSelectedDestinations([]);
    setSelectedOriginUF('');
    setSelectedDestinationUF('');
    setPendingDestinationUnitId(0);
    setOriginUnits([]);
    setDestinationUnits([]);
  };

  const openCreateModal = () => {
    resetForm();
    setModalOpen(true);
  };

  const formatUnit = (unit?: PrisonUnit) =>
    unit ? `${unit.name} (${unit.uf_code})` : 'N/A';

  const formatUnits = (units?: PrisonUnit[]) => {
    if (!units?.length) return 'N/A';
    return units.map((unit) => formatUnit(unit)).join('; ');
  };

  const handleAddDestination = () => {
    if (pendingDestinationUnitId <= 0) {
      setError('Selecione uma unidade prisional de destino para adicionar.');
      return;
    }

    if (formData.destination_unit_ids.includes(pendingDestinationUnitId)) {
      setError('Esta unidade de destino já foi adicionada.');
      return;
    }

    const unit = destinationUnits.find((item) => item.id === pendingDestinationUnitId);
    if (!unit) return;

    setFormData({
      ...formData,
      destination_unit_ids: [...formData.destination_unit_ids, pendingDestinationUnitId]
    });
    setSelectedDestinations([...selectedDestinations, unit]);
    setPendingDestinationUnitId(0);
    setError('');
  };

  const handleRemoveDestination = (unitId: number) => {
    setFormData({
      ...formData,
      destination_unit_ids: formData.destination_unit_ids.filter((id) => id !== unitId)
    });
    setSelectedDestinations(selectedDestinations.filter((unit) => unit.id !== unitId));
  };

  return (
    <Layout>
      <div className="page-container">
        <div className="page-header">
          <div className="page-header-content">
            <h1 className="text-3xl font-bold text-gray-900">Reeducandos</h1>
            <p className="text-gray-600">Gerencie o cadastro de reeducandos</p>
          </div>
          <Button onClick={openCreateModal}>
            + Novo Reeducando
          </Button>
        </div>

        <div className="table-wrapper">
          {loading ? (
            <div className="table-loading">
              <div className="spinner"></div>
              <p className="mt-4">Carregando reeducandos...</p>
            </div>
          ) : inmates.length === 0 ? (
            <div className="table-empty">
              Nenhum reeducando encontrado
            </div>
          ) : (
            <div>
              <table className="table">
                <thead>
                  <tr>
                    <th>ID</th>
                    <th>Unidade origem</th>
                    <th>Unidades destino</th>
                    <th>Regime</th>
                    <th className="text-right">Ações</th>
                  </tr>
                </thead>
                <tbody>
                  {inmates.map((inmate) => (
                    <tr key={inmate.id}>
                      <td>{inmate.id}</td>
                      <td>{formatUnit(inmate.origin_unit)}</td>
                      <td>{formatUnits(inmate.destination_units)}</td>
                      <td>
                        {inmate.custody === 'CLOSED' ? 'Fechado' : inmate.custody === 'SEMI_OPEN' ? 'Semiaberto' : 'Aberto'}
                      </td>
                      <td className="text-right">
                        <div className="table-actions">
                          <Button variant="ghost" size="sm" onClick={() => handleEdit(inmate.id)}>
                            Editar
                          </Button>
                          <Button variant="danger" size="sm" onClick={() => handleDelete(inmate.id)}>
                            Excluir
                          </Button>
                        </div>
                      </td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          )}
        </div>

        <Modal
          isOpen={modalOpen}
          onClose={() => setModalOpen(false)}
          title={editingId ? 'Editar Reeducando' : 'Novo Reeducando'}
          size="lg"
        >
          <form onSubmit={handleSubmit} className="space-y-4">
            <div className="form-section">
              <h3 className="form-section-title">Origem</h3>
              <div className="form-grid-2">
                <Select
                  label="UF Origem"
                  value={selectedOriginUF}
                  onChange={(e) => {
                    setSelectedOriginUF(e.target.value);
                    setFormData({ ...formData, origin_unit_id: 0 });
                    if (e.target.value) {
                      loadOriginUnits(e.target.value);
                    } else {
                      setOriginUnits([]);
                    }
                  }}
                  options={[
                    { value: '', label: 'Selecione a UF' },
                    ...ufs.map(uf => ({ value: uf.code, label: `${uf.code} - ${uf.name}` }))
                  ]}
                  required
                />
                <Select
                  label="Unidade prisional origem"
                  value={formData.origin_unit_id}
                  onChange={(e) => setFormData({ ...formData, origin_unit_id: Number(e.target.value) })}
                  options={[
                    { value: '0', label: 'Selecione a unidade' },
                    ...originUnits.map(unit => ({ value: String(unit.id), label: unit.name }))
                  ]}
                  required
                  disabled={!selectedOriginUF || originUnits.length === 0}
                />
              </div>
            </div>

            <Select
              label="Regime de Custódia"
              value={formData.custody}
              onChange={(e) => setFormData({ ...formData, custody: e.target.value as any })}
              options={[
                { value: 'CLOSED', label: 'Fechado' },
                { value: 'SEMI_OPEN', label: 'Semiaberto' },
                { value: 'OPEN', label: 'Aberto' }
              ]}
            />

            <div className="form-section">
              <h3 className="form-section-title">Destinos</h3>
              <p className="text-sm text-gray-600 mb-3">
                Adicione uma ou mais unidades prisionais de destino para este reeducando.
              </p>
              <div className="form-grid-2">
                <Select
                  label="UF Destino"
                  value={selectedDestinationUF}
                  onChange={(e) => {
                    setSelectedDestinationUF(e.target.value);
                    setPendingDestinationUnitId(0);
                    if (e.target.value) {
                      loadDestinationUnits(e.target.value);
                    } else {
                      setDestinationUnits([]);
                    }
                  }}
                  options={[
                    { value: '', label: 'Selecione a UF' },
                    ...ufs.map(uf => ({ value: uf.code, label: `${uf.code} - ${uf.name}` }))
                  ]}
                />
                <Select
                  label="Unidade prisional destino"
                  value={pendingDestinationUnitId}
                  onChange={(e) => setPendingDestinationUnitId(Number(e.target.value))}
                  options={[
                    { value: '0', label: 'Selecione a unidade' },
                    ...destinationUnits.map(unit => ({ value: String(unit.id), label: unit.name }))
                  ]}
                  disabled={!selectedDestinationUF || destinationUnits.length === 0}
                />
              </div>
              <div className="mt-3">
                <Button type="button" variant="secondary" onClick={handleAddDestination}>
                  + Adicionar destino
                </Button>
              </div>

              {selectedDestinations.length > 0 && (
                <ul className="destination-units-list mt-4">
                  {selectedDestinations.map((unit) => (
                    <li key={unit.id} className="destination-units-item">
                      <span>{formatUnit(unit)}</span>
                      <button
                        type="button"
                        className="destination-units-remove"
                        onClick={() => handleRemoveDestination(unit.id)}
                      >
                        Remover
                      </button>
                    </li>
                  ))}
                </ul>
              )}
            </div>

            <div className="form-section">
              <h3 className="form-section-title">Responsável</h3>
              <div className="form-grid-2">
                <Input
                  label="Advogado"
                  value={formData.responsible.attorney}
                  onChange={(e) => setFormData({ ...formData, responsible: { ...formData.responsible, attorney: e.target.value } })}
                />
                <PhoneInput
                  label="Telefone"
                  value={formData.responsible.phone}
                  onChange={(e) => setFormData({ ...formData, responsible: { ...formData.responsible, phone: e.target.value } })}
                />
              </div>
            </div>

            {error && (
              <div className="bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded-lg text-sm">
                {error}
              </div>
            )}

            <div className="flex justify-end space-x-2 pt-4">
              <Button type="button" variant="secondary" onClick={() => setModalOpen(false)}>
                Cancelar
              </Button>
              <Button type="submit">
                {editingId ? 'Salvar' : 'Criar'}
              </Button>
            </div>
          </form>
        </Modal>
      </div>
    </Layout>
  );
}
