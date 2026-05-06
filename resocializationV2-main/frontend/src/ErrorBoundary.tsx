import { Component } from 'react';
import type { ErrorInfo, ReactNode } from 'react';

interface Props {
  children: ReactNode;
}

interface State {
  hasError: boolean;
  error: Error | null;
}

class ErrorBoundary extends Component<Props, State> {
  public state: State = {
    hasError: false,
    error: null,
  };

  public static getDerivedStateFromError(error: Error): State {
    return { hasError: true, error };
  }

  public componentDidCatch(error: Error, errorInfo: ErrorInfo) {
    console.error('ErrorBoundary caught an error:', error, errorInfo);
  }

  public render() {
    if (this.state.hasError) {
      return (
        <div style={{
          display: 'flex',
          flexDirection: 'column',
          alignItems: 'center',
          justifyContent: 'center',
          height: '100vh',
          padding: '2rem',
        }}>
          <h1>Algo deu errado</h1>
          <p>Houve um erro ao carregar a aplicação.</p>
          <pre style={{
            background: '#f5f5f5',
            padding: '1rem',
            borderRadius: '4px',
            maxWidth: '600px',
            overflow: 'auto'
          }}>
            {this.state.error?.message}
          </pre>
          <button
            onClick={() => window.location.href = '/login'}
            style={{
              marginTop: '1rem',
              padding: '0.5rem 1rem',
              cursor: 'pointer'
            }}
          >
            Ir para Login
          </button>
        </div>
      );
    }

    return this.props.children;
  }
}

export default ErrorBoundary;
