import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { apiListUsers, apiCreateUser, apiUpdateUser, apiDeleteUser } from '../api';
import type { User } from '../types';
import Layout from '../components/Layout';
import Button from '../components/Button';
import Modal from '../components/Modal';
import Input from '../components/Input';

export default function Users() {
  const navigate = useNavigate();
  const [users, setUsers] = useState<User[]>([]);
  const [loading, setLoading] = useState(true);
  const [modalOpen, setModalOpen] = useState(false);
  const [editingUser, setEditingUser] = useState<User | null>(null);
  const [formData, setFormData] = useState({ name: '', email: '', password: '' });
  const [error, setError] = useState('');

  useEffect(() => {
    loadUsers();
  }, []);

  const loadUsers = async () => {
    try {
      setLoading(true);
      const data = await apiListUsers();
      setUsers(data);
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
      if (editingUser) {
        await apiUpdateUser(editingUser.id, formData.name, formData.email, editingUser.status);
      } else {
        await apiCreateUser(formData.name, formData.email, formData.password);
      }
      setModalOpen(false);
      setFormData({ name: '', email: '', password: '' });
      setEditingUser(null);
      loadUsers();
    } catch (err: any) {
      setError(err.message);
    }
  };

  const handleEdit = (user: User) => {
    navigate(`/users/edit/${user.id}`);
  };

  const handleDelete = async (id: number) => {
    if (!confirm('Tem certeza que deseja excluir este usuário?')) return;
    try {
      await apiDeleteUser(id);
      loadUsers();
    } catch (err: any) {
      setError(err.message);
    }
  };

  const openCreateModal = () => {
    setEditingUser(null);
    setFormData({ name: '', email: '', password: '' });
    setModalOpen(true);
  };

  return (
    <Layout>
      <div className="page-container">
        <div className="page-header">
          <div className="page-header-content">
            <h1 className="text-3xl font-bold text-gray-900">Usuários</h1>
            <p className="text-gray-600">Gerencie os usuários do sistema</p>
          </div>
          <Button onClick={openCreateModal}>
            + Novo Usuário
          </Button>
        </div>

        <div className="table-wrapper">
          {loading ? (
            <div className="table-loading">
              <div className="spinner"></div>
              <p className="mt-4">Carregando usuários...</p>
            </div>
          ) : users.length === 0 ? (
            <div className="table-empty">
              Nenhum usuário encontrado
            </div>
          ) : (
            <div>
              <table className="table">
                <thead>
                  <tr>
                    <th>ID</th>
                    <th>Nome</th>
                    <th>Email</th>
                    <th>Status</th>
                    <th>Role</th>
                    <th className="text-right">Ações</th>
                  </tr>
                </thead>
                <tbody>
                  {users.map((user) => (
                    <tr key={user.id}>
                      <td>{user.id}</td>
                      <td>{user.name}</td>
                      <td className="text-gray-600">{user.email}</td>
                      <td>
                        <span style={{
                          padding: '0.25rem 0.5rem',
                          borderRadius: '4px',
                          fontSize: '12px',
                          fontWeight: '500',
                          backgroundColor: user.status === 'ativo' ? '#d4edda' : '#f8d7da',
                          color: user.status === 'ativo' ? '#155724' : '#721c24'
                        }}>
                          {user.status === 'ativo' ? 'Ativo' : 'Inativo'}
                        </span>
                      </td>
                      <td className="text-gray-600">{user.role}</td>
                      <td className="text-right">
                        <div className="table-actions">
                          <Button variant="ghost" size="sm" onClick={() => handleEdit(user)}>
                            Editar
                          </Button>
                          <Button variant="danger" size="sm" onClick={() => handleDelete(user.id)}>
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
          title={editingUser ? 'Editar Usuário' : 'Novo Usuário'}
        >
          <form onSubmit={handleSubmit} className="space-y-4">
            <Input
              label="Nome"
              value={formData.name}
              onChange={(e) => setFormData({ ...formData, name: e.target.value })}
              required
            />
            <Input
              label="Email"
              type="email"
              value={formData.email}
              onChange={(e) => setFormData({ ...formData, email: e.target.value })}
              required
            />
            {!editingUser && (
              <Input
                label="Senha"
                type="password"
                value={formData.password}
                onChange={(e) => setFormData({ ...formData, password: e.target.value })}
                required
              />
            )}
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
                {editingUser ? 'Salvar' : 'Criar'}
              </Button>
            </div>
          </form>
        </Modal>
      </div>
    </Layout>
  );
}
