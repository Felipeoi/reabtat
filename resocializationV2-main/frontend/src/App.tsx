import { Routes, Route, Navigate } from 'react-router-dom';
import Login from './pages/Login';
import Signup from './pages/Signup';
import Users from './pages/Users';
import EditUser from './pages/EditUser';
import Inmates from './pages/Inmates';
import Matches from './pages/Matches';
import MatchDetails from './pages/MatchDetails';
import NotFound from './pages/NotFound';
import { isAuthed } from './auth';
import { useEffect, useState } from 'react';
import { apiMe } from './api';

// Componente ProtectedRoute SIMPLES
function ProtectedRoute({ children, requiredRole }: { children: React.ReactNode; requiredRole?: string }) {
  const [user, setUser] = useState<any>(null);
  const [loading, setLoading] = useState(requiredRole ? true : false);
  const [checked, setChecked] = useState(false);

  useEffect(() => {
    // Se não precisa verificar role, não carrega nada
    if (!requiredRole) {
      setChecked(true);
      return;
    }

    // Se precisa verificar role, busca dados do usuário
    if (isAuthed() && !checked) {
      apiMe()
        .then(data => {
          setUser(data);
          setLoading(false);
          setChecked(true);
        })
        .catch(err => {
          console.error('Erro ao carregar usuário:', err);
          setLoading(false);
          setChecked(true);
        });
    }
  }, [requiredRole, checked]);

  if (!isAuthed()) {
    return <Navigate to="/login" replace />;
  }

  // Se requer role e ainda está carregando
  if (requiredRole && loading) {
    return <div style={{ display: 'flex', justifyContent: 'center', alignItems: 'center', height: '100vh' }}>Carregando...</div>;
  }

  // Se requer role e usuário não tem
  if (requiredRole && user && user.role !== requiredRole) {
    return (
      <div style={{ display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center', height: '100vh', gap: '1rem' }}>
        <h2>⛔ Acesso Negado</h2>
        <p>Você não tem permissão para acessar esta página.</p>
        <p>Necessário: <strong>{requiredRole}</strong></p>
        <p>Seu perfil: <strong>{user.role}</strong></p>
        <button onClick={() => window.location.href = '/inmates'} style={{ padding: '0.5rem 1rem', cursor: 'pointer' }}>
          Voltar para reeducandos
        </button>
      </div>
    );
  }

  return <>{children}</>;
}

export default function App() {
  return (
    <Routes>
      <Route path="/" element={<Navigate to="/login" replace />} />
      <Route path="/login" element={<Login />} />
      <Route path="/signup" element={<Signup />} />

      <Route
        path="/users"
        element={
          <ProtectedRoute requiredRole="admin">
            <Users />
          </ProtectedRoute>
        }
      />

      <Route
        path="/users/edit/:id"
        element={
          <ProtectedRoute requiredRole="admin">
            <EditUser />
          </ProtectedRoute>
        }
      />

      <Route
        path="/inmates"
        element={
          <ProtectedRoute>
            <Inmates />
          </ProtectedRoute>
        }
      />

      <Route
        path="/matches"
        element={
          <ProtectedRoute>
            <Matches />
          </ProtectedRoute>
        }
      />

      <Route
        path="/matches/:myInmateId/:matchedInmateId"
        element={
          <ProtectedRoute>
            <MatchDetails />
          </ProtectedRoute>
        }
      />

      <Route path="*" element={<NotFound />} />
    </Routes>
  );
}
