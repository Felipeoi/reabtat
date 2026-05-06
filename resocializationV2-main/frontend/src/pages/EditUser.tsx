import { useState, useEffect } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import { apiGetUser, apiUpdateUser } from '../api';
import Layout from '../components/Layout';
import '../styles/editUser.css';

export default function EditUser() {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [status, setStatus] = useState('ativo');
  const [password, setPassword] = useState('');
  const [loading, setLoading] = useState(true);
  const [saving, setSaving] = useState(false);
  const [error, setError] = useState('');

  useEffect(() => {
    if (!id) return;

    apiGetUser(parseInt(id))
      .then(user => {
        setName(user.name);
        setEmail(user.email);
        setStatus(user.status);
        setLoading(false);
      })
      .catch(err => {
        console.error('Erro ao carregar usuário:', err);
        setError('Erro ao carregar dados do usuário');
        setLoading(false);
      });
  }, [id]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError('');
    setSaving(true);

    try {
      await apiUpdateUser(parseInt(id!), name, email, status, password || undefined);
      navigate('/users');
    } catch (err: any) {
      setError(err.message || 'Erro ao atualizar usuário');
    } finally {
      setSaving(false);
    }
  };

  if (loading) {
    return (
      <Layout>
        <div className="loading-container">
          <div className="loading-spinner"></div>
          <p className="loading-text">Carregando dados do usuário...</p>
        </div>
      </Layout>
    );
  }

  return (
    <Layout>
      <div className="edit-user-container">
        <div className="edit-user-card">
          <div className="edit-user-header">
            <h1 className="edit-user-title">Editar Usuário</h1>
            <p className="edit-user-subtitle">Atualize as informações do usuário no sistema</p>
          </div>

          <form onSubmit={handleSubmit} className="edit-user-form">
            <div className="form-group">
              <label className="form-label" htmlFor="name">Nome Completo</label>
              <input
                id="name"
                type="text"
                className="form-input"
                value={name}
                onChange={(e) => setName(e.target.value)}
                placeholder="Digite o nome completo"
                required
              />
            </div>

            <div className="form-group">
              <label className="form-label" htmlFor="email">E-mail</label>
              <input
                id="email"
                type="email"
                className="form-input"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                placeholder="Digite o e-mail"
                required
              />
            </div>

            <div className="form-group">
              <label className="form-label" htmlFor="status">Status do Usuário</label>
              <select
                id="status"
                className="form-select"
                value={status}
                onChange={(e) => setStatus(e.target.value)}
              >
                <option value="ativo">✓ Ativo - Pode acessar o sistema</option>
                <option value="inativo">✗ Inativo - Bloqueado no sistema</option>
              </select>
              <span className="form-hint">
                Usuários inativos não conseguem fazer login
              </span>
            </div>

            <div className="form-group">
              <label className="form-label" htmlFor="password">Nova Senha</label>
              <input
                id="password"
                type="password"
                className="form-input"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                placeholder="••••••••"
              />
              <span className="form-hint">
                Deixe em branco para manter a senha atual
              </span>
            </div>

            {error && (
              <div className="form-error">
                ⚠️ {error}
              </div>
            )}

            <div className="form-actions">
              <button type="submit" className="btn btn-primary" disabled={saving}>
                {saving ? '⏳ Salvando...' : '💾 Salvar Alterações'}
              </button>
              <button type="button" className="btn btn-secondary" onClick={() => navigate('/users')}>
                ← Cancelar
              </button>
            </div>
          </form>
        </div>
      </div>
    </Layout>
  );
}
