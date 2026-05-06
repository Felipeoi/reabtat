import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { apiFindMatches } from '../api';
import type { MatchResult } from '../types';
import Layout from '../components/Layout';
import Button from '../components/Button';
import '../styles/matches.css';

export default function Matches() {
  const [matches, setMatches] = useState<MatchResult[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  const navigate = useNavigate();

  useEffect(() => {
    loadMatches();
  }, []);

  const loadMatches = async () => {
    try {
      setLoading(true);
      setError('');
      const data = await apiFindMatches();
      // Garantir que sempre seja um array, mesmo se a API retornar null/undefined
      setMatches(Array.isArray(data) ? data : []);
    } catch (err: any) {
      setError(err.message || 'Erro ao buscar matches');
      setMatches([]); // Garantir estado consistente em caso de erro
    } finally {
      setLoading(false);
    }
  };

  const getCustodyLabel = (custody: string) => {
    switch (custody) {
      case 'CLOSED': return 'Fechado';
      case 'SEMI_OPEN': return 'Semi-Aberto';
      case 'OPEN': return 'Aberto';
      default: return custody;
    }
  };

  const handleViewDetails = (myInmateId: number, matchedInmateId: number) => {
    navigate(`/matches/${myInmateId}/${matchedInmateId}`);
  };

  return (
    <Layout>
      <div className="page-container">
        <div className="page-header">
          <div className="page-header-content">
            <h1 className="text-3xl font-bold text-gray-900">Matches Disponíveis</h1>
            <p className="text-gray-600">Encontre reeducandos compatíveis para troca de localização</p>
          </div>
          <Button onClick={loadMatches} variant="secondary">
            Atualizar
          </Button>
        </div>

        {error && (
          <div className="bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded-lg mb-4">
            {error}
          </div>
        )}

        <div>
          {loading ? (
            <div className="matches-loading">
              <div className="spinner"></div>
              <p className="matches-loading-text">Buscando matches perfeitos...</p>
            </div>
          ) : !matches || matches.length === 0 ? (
            <div className="matches-empty">
              <div className="matches-empty-card">
                <h3 className="matches-empty-title">Sem matches por enquanto</h3>
                <p className="matches-empty-text">
                  Ainda não há matches disponíveis. Continue cadastrando reeducandos e em breve encontraremos compatibilidades para trocas!
                </p>
              </div>
            </div>
          ) : (
            <div className="matches-grid">
              {matches.map((match, index) => {
                // Proteção contra dados faltantes
                if (!match || !match.my_inmate || !match.matched_inmate) {
                  return null;
                }

                return (
                  <div
                    key={index}
                    onClick={() => handleViewDetails(match.my_inmate.id, match.matched_inmate.id)}
                    className="match-card"
                  >
                    {/* Header */}
                    <div className="match-card-header">
                      <div className="match-card-badges">
                        <span className="match-badge match-badge-perfect">
                          Match Perfeito
                        </span>
                        <span className={`match-badge ${
                          match.custody === 'CLOSED' ? 'match-badge-custody-closed' :
                          match.custody === 'SEMI_OPEN' ? 'match-badge-custody-semiopen' :
                          'match-badge-custody-open'
                        }`}>
                          {getCustodyLabel(match.custody)}
                        </span>
                      </div>
                      <span className="match-badge match-badge-score">
                        {match.match_score || 0}%
                      </span>
                    </div>

                    {/* Conteúdo */}
                    <div className="match-card-content">
                      {/* Seu Reeducando */}
                      <div className="match-inmate-section blue">
                        <div className="match-inmate-header">
                          <span className="match-inmate-title blue">Seu Reeducando #{match.my_inmate.id}</span>
                        </div>
                        <div className="match-inmate-info">
                          <div className="match-info-row">
                            <span className="match-info-label">Origem:</span>
                            <span className="match-info-value">
                              {match.my_inmate.origin?.name || 'N/A'}, {match.my_inmate.origin?.uf_code || 'N/A'}
                            </span>
                          </div>
                          <div className="match-info-row">
                            <span className="match-info-label">Destino:</span>
                            <span className="match-info-value">
                              {match.my_inmate.destination?.name || 'N/A'}, {match.my_inmate.destination?.uf_code || 'N/A'}
                            </span>
                          </div>
                        </div>
                      </div>

                      {/* Separador */}
                      <div className="match-separator">
                        <span className="match-separator-badge">↕ Troca</span>
                      </div>

                      {/* Match Encontrado */}
                      <div className="match-inmate-section green">
                        <div className="match-inmate-header">
                          <span className="match-inmate-title green">Match Encontrado #{match.matched_inmate.id}</span>
                        </div>
                        <div className="match-inmate-info">
                          <div className="match-info-row">
                            <span className="match-info-label">Origem:</span>
                            <span className="match-info-value">
                              {match.matched_inmate.origin?.name || 'N/A'}, {match.matched_inmate.origin?.uf_code || 'N/A'}
                            </span>
                          </div>
                          <div className="match-info-row">
                            <span className="match-info-label">Destino:</span>
                            <span className="match-info-value">
                              {match.matched_inmate.destination?.name || 'N/A'}, {match.matched_inmate.destination?.uf_code || 'N/A'}
                            </span>
                          </div>
                        </div>
                      </div>
                    </div>

                    {/* Footer */}
                    <div className="match-card-footer">
                      <span className="match-card-footer-text">Clique para ver mais detalhes →</span>
                    </div>
                  </div>
                );
              })}
            </div>
          )}
        </div>
      </div>
    </Layout>
  );
}
