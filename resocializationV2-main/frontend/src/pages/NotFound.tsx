import { Link } from 'react-router-dom';
import Button from '../components/Button';

export default function NotFound() {
  return (
    <div className="not-found">
      <div className="not-found-container">
        <div className="not-found-code">404</div>
        <div className="not-found-card">
          <h1 className="not-found-title">Página não encontrada</h1>
          <p className="not-found-message">
            A página que você está procurando não existe ou foi removida.
          </p>
          <Link to="/users">
            <Button size="lg">🏠 Voltar para o início</Button>
          </Link>
        </div>
      </div>
    </div>
  );
}
