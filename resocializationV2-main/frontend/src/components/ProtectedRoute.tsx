import type { ReactNode } from 'react';
import { useEffect } from 'react';
import { Navigate, useLocation } from 'react-router-dom';
import { isAuthed } from '../auth';
import { useUser } from '../userContext';

interface ProtectedRouteProps {
  children: ReactNode;
  requiredRole?: string;
}

export default function ProtectedRoute({ children, requiredRole }: ProtectedRouteProps) {
  const loc = useLocation();
  const { user, loading, refreshUser } = useUser();

  // Se não está autenticado, redireciona para login imediatamente
  if (!isAuthed()) {
    return <Navigate to="/login" replace state={{ from: loc }} />;
  }

  // Carrega dados do usuário se não tiver (lazy loading)
  useEffect(() => {
    if (isAuthed() && !user && !loading) {
      refreshUser();
    }
  }, [user, loading, refreshUser]);

  // Aguarda carregar dados do usuário se houver requiredRole
  if (requiredRole && (loading || !user)) {
    return (
      <div style={{
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        height: '100vh',
        fontSize: '18px',
        color: '#666'
      }}>
        Carregando...
      </div>
    );
  }

  // Se requer role específica e usuário não tem
  if (requiredRole && user && user.role !== requiredRole) {
    return (
      <div style={{
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        height: '100vh',
        flexDirection: 'column',
        gap: '1rem',
        padding: '2rem',
        textAlign: 'center'
      }}>
        <h2 style={{ fontSize: '2rem', marginBottom: '1rem' }}>⛔ Acesso Negado</h2>
        <p>Você não tem permissão para acessar esta página.</p>
        <p>Necessário: <strong>{requiredRole}</strong></p>
        <p>Seu perfil: <strong>{user.role}</strong></p>
        <button
          onClick={() => window.location.href = '/inmates'}
          style={{
            padding: '0.75rem 1.5rem',
            marginTop: '1rem',
            cursor: 'pointer',
            backgroundColor: '#3b82f6',
            color: 'white',
            border: 'none',
            borderRadius: '4px',
            fontSize: '16px'
          }}
        >
          Voltar para Inmates
        </button>
      </div>
    );
  }

  return <>{children}</>;
}
