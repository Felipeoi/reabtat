import { createContext, useContext, useState, useCallback, useEffect } from 'react';
import type { ReactNode } from 'react';
import type { User } from './types';
import { apiMe } from './api';
import { isAuthed } from './auth';

interface UserContextType {
  user: User | null;
  loading: boolean;
  refreshUser: () => Promise<void>;
  hasRole: (role: string) => boolean;
}

const UserContext = createContext<UserContextType>({
  user: null,
  loading: false,
  refreshUser: async () => {},
  hasRole: () => false,
});

export function UserProvider({ children }: { children: ReactNode }) {
  const [user, setUser] = useState<User | null>(null);
  const [loading, setLoading] = useState(false);
  const [initialized, setInitialized] = useState(false);

  const refreshUser = useCallback(async () => {
    // Se não está autenticado, não precisa buscar dados
    if (!isAuthed()) {
      setUser(null);
      setLoading(false);
      setInitialized(true);
      return;
    }

    setLoading(true);
    try {
      const userData = await apiMe();
      setUser(userData);
    } catch (error) {
      console.error('Error fetching user data:', error);
      setUser(null);
      // Se der erro de autenticação, limpa o token
      if (error && typeof error === 'object' && 'status' in error && error.status === 401) {
        localStorage.removeItem('auth_token');
      }
    } finally {
      setLoading(false);
      setInitialized(true);
    }
  }, []);

  // Carrega dados do usuário na primeira renderização SE estiver autenticado
  useEffect(() => {
    if (!initialized && isAuthed()) {
      refreshUser();
    } else if (!initialized) {
      setInitialized(true);
    }
  }, [initialized, refreshUser]);

  const hasRole = useCallback((role: string): boolean => {
    return user?.role === role;
  }, [user]);

  const value = {
    user,
    loading,
    refreshUser,
    hasRole,
  };

  return (
    <UserContext.Provider value={value}>
      {children}
    </UserContext.Provider>
  );
}

export function useUser() {
  const context = useContext(UserContext);
  return context;
}
