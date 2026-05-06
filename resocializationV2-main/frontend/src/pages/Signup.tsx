import { useState } from 'react';
import type { FormEvent } from 'react';
import { Link } from 'react-router-dom';
import { apiSignup, apiLogin } from '../api';
import { setToken } from '../auth';
import Button from '../components/Button';
import Input from '../components/Input';

export default function Signup() {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [telefone, setTelefone] = useState('');
  const [error, setError] = useState('');
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault();
    setError('');
    setLoading(true);

    try {
      await apiSignup(name, email, password, telefone);
      const res = await apiLogin(email, password);
      setToken(res.token);
      // Recarrega a página para que o UserContext carregue os dados do usuário
      window.location.href = '/inmates';
    } catch (err: any) {
      setError(err.message || 'Erro ao criar conta');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="auth-page">
      <div className="auth-container">
        <div className="auth-logo">R</div>

        <h2 className="auth-title">Criar nova conta</h2>
        <p className="auth-subtitle">Sistema de Ressocialização</p>

        <div className="auth-card">
          <form className="auth-form" onSubmit={handleSubmit}>
            <Input
              label="Nome"
              type="text"
              value={name}
              onChange={(e) => setName(e.target.value)}
              placeholder="Seu nome completo"
              required
            />

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
              label="Telefone (opcional)"
              type="tel"
              value={telefone}
              onChange={(e) => setTelefone(e.target.value)}
              placeholder="+5511999999999"
              autoComplete="tel"
            />

            <Input
              label="Senha"
              type="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              placeholder="••••••••"
              required
              autoComplete="new-password"
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
              {loading ? 'Criando conta...' : 'Criar conta'}
            </Button>
          </form>

          <div className="auth-divider">
            Já tem uma conta?{' '}
            <Link to="/login" className="auth-link">
              Entrar
            </Link>
          </div>
        </div>
      </div>
    </div>
  );
}
