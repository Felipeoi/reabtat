import { useState } from 'react';
import type { FormEvent } from 'react';
import { Link } from 'react-router-dom';
import { apiLogin } from '../api';
import { setToken } from '../auth';
import Button from '../components/Button';
import Input from '../components/Input';

export default function Login() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault();
    setError('');
    setLoading(true);

    try {
      const res = await apiLogin(email.trim(), password);
      setToken(res.token);
      // Recarrega a página para que o UserContext carregue os dados do usuário
      window.location.href = '/inmates';
    } catch (err: any) {
      setError(err.message || 'Erro ao fazer login');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="auth-page">
      <div className="auth-container">
        <img
          src="/rehabtat-logo.png"
          alt="Rehab TaT — Reabilitação Transferência a Transferência"
          className="auth-brand-logo"
        />

        <h2 className="auth-title">Bem-vindo</h2>
        <p className="auth-subtitle">Sistema de Ressocialização</p>

        <div className="auth-card">
          <form className="auth-form" onSubmit={handleSubmit}>
            <Input
              label="Email"
              type="email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              placeholder="seu@email.com"
              required
              autoComplete="email"
            />

            <Input
              label="Senha"
              type="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              placeholder="••••••••"
              required
              autoComplete="current-password"
            />

            {error && (
              <div className="auth-error">
                {error}
              </div>
            )}

            <Button
              type="submit"
              disabled={loading}
              className="w-full mt-4"
              size="lg"
            >
              {loading ? 'Entrando...' : 'Entrar'}
            </Button>
          </form>

          <div className="auth-divider">
            Não tem uma conta?{' '}
            <Link to="/signup" className="auth-link">
              Criar conta
            </Link>
          </div>
        </div>
      </div>
    </div>
  );
}
