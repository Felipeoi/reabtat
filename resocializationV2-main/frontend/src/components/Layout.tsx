import type { ReactNode } from 'react';
import { Link, useNavigate, useLocation } from 'react-router-dom';
import { isAuthed, logout } from '../auth';

interface LayoutProps {
  children: ReactNode;
}

export default function Layout({ children }: LayoutProps) {
  const navigate = useNavigate();
  const location = useLocation();
  const authenticated = isAuthed();

  const handleLogout = () => {
    logout();
    navigate('/login');
  };

  const navLinks = [
    { path: '/users', label: 'Usuários' },
    { path: '/inmates', label: 'Reeducandos' },
    { path: '/matches', label: 'Correlações' },
  ];

  const isActive = (path: string) => location.pathname.startsWith(path);

  return (
    <div>
      {authenticated && (
        <nav className="navbar">
          <div className="navbar-container">
            <Link to="/inmates" className="navbar-brand">
              <img
                src="/rehabtat-logo.png"
                alt="Rehab TaT — Reabilitação Transferência a Transferência"
                className="navbar-logo"
              />
            </Link>

            <div className="navbar-links">
              {navLinks.map((link) => (
                <Link
                  key={link.path}
                  to={link.path}
                  className={`navbar-link ${isActive(link.path) ? 'active' : ''}`.trim()}
                >
                  {link.label}
                </Link>
              ))}

              <button onClick={handleLogout} className="navbar-logout">
                Sair
              </button>
            </div>
          </div>
        </nav>
      )}

      <main>
        {children}
      </main>
    </div>
  );
}
