import { useState, useEffect } from 'react';
import { apiListInmates, apiCreateInmate, apiUpdateInmate, apiDeleteInmate, apiGetInmate, apiListUFs, apiListCities } from '../api';
import type { InmatesList, UF, City } from '../types';
import Layout from '../components/Layout';
import Button from '../components/Button';
import Modal from '../components/Modal';
import Input from '../components/Input';
import Select from '../components/Select';
import PhoneInput from '../components/PhoneInput';

export default function Inmates() {
  const [inmates, setInmates] = useState<InmatesList[]>([]);
  const [ufs, setUfs] = useState<UF[]>([]);
  const [originCities, setOriginCities] = useState<City[]>([]);
  const [destinationCities, setDestinationCities] = useState<City[]>([]);
  const [loading, setLoading] = useState(true);
  const [modalOpen, setModalOpen] = useState(false);
  const [editingId, setEditingId] = useState<number | null>(null);
  const [formData, setFormData] = useState({
    origin_id: 0,
    custody: 'CLOSED' as 'CLOSED' | 'SEMI_OPEN' | 'OPEN',
    destination_id: 0,
    responsible: { attorney: '', phone: '' }
  });
  const [selectedOriginUF, setSelectedOriginUF] = useState('');
  const [selectedDestinationUF, setSelectedDestinationUF] = useState('');
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

  const loadOriginCities = async (ufCode: string) => {
    try {
      const data = await apiListCities({ uf_code: ufCode });
      setOriginCities(data ?? []);
    } catch (err: any) {
      console.error('Erro ao carregar cidades de origem:', err);
    }
  };

  const loadDestinationCities = async (ufCode: string) => {
    try {
      const data = await apiListCities({ uf_code: ufCode });
      setDestinationCities(data ?? []);
    } catch (err: any) {
      console.error('Erro ao carregar cidades de destino:', err);
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
      setEditingId(id);
      setFormData({
        origin_id: inmate.origin_id,
        custody: inmate.custody,
        destination_id: inmate.destination_id,
        responsible: inmate.responsible
      });

      // Load cities for the selected UFs
      if (inmate.origin?.uf_code) {
        setSelectedOriginUF(inmate.origin.uf_code);
        await loadOriginCities(inmate.origin.uf_code);
      }
      if (inmate.destination?.uf_code) {
        setSelectedDestinationUF(inmate.destination.uf_code);
        await loadDestinationCities(inmate.destination.uf_code);
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
      origin_id: 0,
      custody: 'CLOSED',
      destination_id: 0,
      responsible: { attorney: '', phone: '' }
    });
    setSelectedOriginUF('');
    setSelectedDestinationUF('');
    setOriginCities([]);
    setDestinationCities([]);
  };

  const openCreateModal = () => {
    resetForm();
    setModalOpen(true);
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
                    <th>Origem</th>
                    <th>Destino</th>
                    <th>Regime</th>
                    <th className="text-right">Ações</th>
                  </tr>
                </thead>
                <tbody>
                  {inmates.map((inmate) => (
                    <tr key={inmate.id}>
                      <td>{inmate.id}</td>
                      <td>
                        {inmate.origin ? `${inmate.origin.name} - ${inmate.origin.uf_code}` : 'N/A'}
                      </td>
                      <td>
                        {inmate.destination ? `${inmate.destination.name} - ${inmate.destination.uf_code}` : 'N/A'}
                      </td>
                      <td>
                        {inmate.custody === 'CLOSED' ? 'Fechado' : inmate.custody === 'SEMI_OPEN' ? 'Semi-Aberto' : 'Aberto'}
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
                    setFormData({ ...formData, origin_id: 0 });
                    if (e.target.value) {
                      loadOriginCities(e.target.value);
                    } else {
                      setOriginCities([]);
                    }
                  }}
                  options={[
                    { value: '', label: 'Selecione a UF' },
                    ...ufs.map(uf => ({ value: uf.code, label: `${uf.code} - ${uf.name}` }))
                  ]}
                  required
                />
                <Select
                  label="Cidade Origem"
                  value={formData.origin_id}
                  onChange={(e) => setFormData({ ...formData, origin_id: Number(e.target.value) })}
                  options={[
                    { value: '0', label: 'Selecione a cidade' },
                    ...originCities.map(city => ({ value: String(city.id), label: city.name }))
                  ]}
                  required
                  disabled={!selectedOriginUF || originCities.length === 0}
                />
              </div>
            </div>

            <Select
              label="Regime de Custódia"
              value={formData.custody}
              onChange={(e) => setFormData({ ...formData, custody: e.target.value as any })}
              options={[
                { value: 'CLOSED', label: 'Fechado' },
                { value: 'SEMI_OPEN', label: 'Semi-Aberto' },
                { value: 'OPEN', label: 'Aberto' }
              ]}
            />

            <div className="form-section">
              <h3 className="form-section-title">Destino</h3>
              <div className="form-grid-2">
                <Select
                  label="UF Destino"
                  value={selectedDestinationUF}
                  onChange={(e) => {
                    setSelectedDestinationUF(e.target.value);
                    setFormData({ ...formData, destination_id: 0 });
                    if (e.target.value) {
                      loadDestinationCities(e.target.value);
                    } else {
                      setDestinationCities([]);
                    }
                  }}
                  options={[
                    { value: '', label: 'Selecione a UF' },
                    ...ufs.map(uf => ({ value: uf.code, label: `${uf.code} - ${uf.name}` }))
                  ]}
                  required
                />
                <Select
                  label="Cidade Destino"
                  value={formData.destination_id}
                  onChange={(e) => setFormData({ ...formData, destination_id: Number(e.target.value) })}
                  options={[
                    { value: '0', label: 'Selecione a cidade' },
                    ...destinationCities.map(city => ({ value: String(city.id), label: city.name }))
                  ]}
                  required
                  disabled={!selectedDestinationUF || destinationCities.length === 0}
                />
              </div>
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
