import { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { apiFindMatchById } from '../api';
import type { MatchResult } from '../types';
import Layout from '../components/Layout';
import Button from '../components/Button';
import '../styles/match-details.css';

export default function MatchDetails() {
  const { myInmateId, matchedInmateId } = useParams<{ myInmateId: string; matchedInmateId: string }>();
  const [match, setMatch] = useState<MatchResult | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  const navigate = useNavigate();

  useEffect(() => {
    loadMatch();
  }, [myInmateId, matchedInmateId]);

  const loadMatch = async () => {
    if (!myInmateId || !matchedInmateId) {
      setError('IDs inválidos');
      setLoading(false);
      return;
    }

    try {
      setLoading(true);
      setError('');
      const data = await apiFindMatchById(parseInt(myInmateId), parseInt(matchedInmateId));
      setMatch(data || null);
    } catch (err: any) {
      setError(err.message || 'Erro ao carregar detalhes do match');
      setMatch(null);
    } finally {
      setLoading(false);
    }
  };

  const getCustodyLabel = (custody: string) => {
    switch (custody) {
      case 'CLOSED': return 'Fechado';
      case 'SEMI_OPEN': return 'Semiaberto';
      case 'OPEN': return 'Aberto';
      default: return custody;
    }
  };

  const getCustodyClass = (custody: string) => {
    switch (custody) {
      case 'CLOSED': return 'closed';
      case 'SEMI_OPEN': return 'semiopen';
      case 'OPEN': return 'open';
      default: return 'closed';
    }
  };

  if (loading) {
    return (
      <Layout>
        <div className="match-details-loading">
          <div className="spinner"></div>
          <p className="match-details-loading-text">Carregando detalhes do match...</p>
        </div>
      </Layout>
    );
  }

  if (error || !match) {
    return (
      <Layout>
        <div className="match-details-error">
          <div className="match-details-error-card">
            <div className="match-details-error-icon">
              ⚠
            </div>
            <h3 className="match-details-error-title">Erro ao carregar match</h3>
            <p className="match-details-error-text">{error || 'Match não encontrado'}</p>
            <Button onClick={() => navigate('/matches')}>
              Voltar para Matches
            </Button>
          </div>
        </div>
      </Layout>
    );
  }

  return (
    <Layout>
      <div className="page-container">
        {/* Back Button */}
        <div className="match-details-back">
          <Button
            onClick={() => navigate('/matches')}
            variant="secondary"
          >
            ← Voltar
          </Button>
        </div>

        {/* Header */}
        <div className="match-details-header">
          <div className="match-details-header-content">
            <div className="match-details-header-info">
              <h1>Detalhes do Match</h1>
              <p>Análise completa da compatibilidade entre reeducandos</p>
            </div>
            <div className="match-details-header-badges">
              <span className="match-details-badge match-details-badge-perfect">
                ✓ Match Perfeito
              </span>
              <span className={`match-details-badge match-details-badge-custody ${getCustodyClass(match.custody)}`}>
                {getCustodyLabel(match.custody)}
              </span>
              <span className="match-details-badge match-details-badge-score">
                Score: {match.match_score}%
              </span>
            </div>
          </div>
        </div>

        {/* Main Content Grid */}
        <div className="match-details-grid">
          {/* Seu Reeducando */}
          <div className="match-details-inmate-card">
            <div className="match-details-inmate-header">
              <span className="match-details-inmate-id blue">
                {match.my_inmate.id}
              </span>
              <div className="match-details-inmate-title">
                <h2>Seu Reeducando</h2>
                <p>Informações completas do seu interno</p>
              </div>
            </div>

            {/* Localização */}
            <div className="match-details-section">
              <h3 className="match-details-section-title blue">
                📍 Localização
              </h3>
              <div className="match-details-location-card">
                <p className="match-details-location-type">Origem</p>
                <p className="match-details-location-name">{match.my_inmate.origin?.name}</p>
                <p className="match-details-location-state">{match.my_inmate.origin?.uf_code}</p>
                <p className="match-details-location-code">IBGE: {match.my_inmate.origin?.ibge_code}</p>
              </div>
              <div className="match-details-location-arrow">↓</div>
              <div className="match-details-location-card">
                <p className="match-details-location-type">Destino Desejado</p>
                <p className="match-details-location-name">{match.my_inmate.destination?.name}</p>
                <p className="match-details-location-state">{match.my_inmate.destination?.uf_code}</p>
                <p className="match-details-location-code">IBGE: {match.my_inmate.destination?.ibge_code}</p>
              </div>
            </div>

            {/* Responsável */}
            {(match.my_inmate.responsible.attorney || match.my_inmate.responsible.phone) && (
              <div className="match-details-section">
                <h3 className="match-details-section-title blue">
                  👤 Responsável
                </h3>
                {match.my_inmate.responsible.attorney && (
                  <div className="match-details-contact-card">
                    <div>
                      <p className="match-details-contact-label">Advogado</p>
                      <p className="match-details-contact-value">{match.my_inmate.responsible.attorney}</p>
                    </div>
                  </div>
                )}
                {match.my_inmate.responsible.phone && (
                  <div className="match-details-contact-card">
                    <div>
                      <p className="match-details-contact-label">Telefone</p>
                      <p className="match-details-contact-value">{match.my_inmate.responsible.phone}</p>
                    </div>
                  </div>
                )}
              </div>
            )}

            {/* Regime */}
            <div className="match-details-custody-card">
              <div>
                <p className="match-details-custody-label">Regime de Custódia</p>
                <p className="match-details-custody-value">{getCustodyLabel(match.my_inmate.custody)}</p>
              </div>
              <span className={`match-details-custody-badge ${getCustodyClass(match.my_inmate.custody)}`}>
                {match.my_inmate.custody}
              </span>
            </div>
          </div>

          {/* Match Encontrado */}
          <div className="match-details-inmate-card">
            <div className="match-details-inmate-header">
              <span className="match-details-inmate-id green">
                {match.matched_inmate.id}
              </span>
              <div className="match-details-inmate-title">
                <h2>Match Encontrado</h2>
                <p>Reeducando compatível para troca</p>
              </div>
            </div>

            {/* Localização */}
            <div className="match-details-section">
              <h3 className="match-details-section-title green">
                📍 Localização
              </h3>
              <div className="match-details-location-card">
                <p className="match-details-location-type">Origem</p>
                <p className="match-details-location-name">{match.matched_inmate.origin?.name}</p>
                <p className="match-details-location-state">{match.matched_inmate.origin?.uf_code}</p>
                <p className="match-details-location-code">IBGE: {match.matched_inmate.origin?.ibge_code}</p>
              </div>
              <div className="match-details-location-arrow">↓</div>
              <div className="match-details-location-card">
                <p className="match-details-location-type">Destino Desejado</p>
                <p className="match-details-location-name">{match.matched_inmate.destination?.name}</p>
                <p className="match-details-location-state">{match.matched_inmate.destination?.uf_code}</p>
                <p className="match-details-location-code">IBGE: {match.matched_inmate.destination?.ibge_code}</p>
              </div>
            </div>

            {/* Responsável */}
            {(match.matched_inmate.responsible.attorney || match.matched_inmate.responsible.phone) && (
              <div className="match-details-section">
                <h3 className="match-details-section-title green">
                  👤 Responsável
                </h3>
                {match.matched_inmate.responsible.attorney && (
                  <div className="match-details-contact-card">
                    <div>
                      <p className="match-details-contact-label">Advogado</p>
                      <p className="match-details-contact-value">{match.matched_inmate.responsible.attorney}</p>
                    </div>
                  </div>
                )}
                {match.matched_inmate.responsible.phone && (
                  <div className="match-details-contact-card">
                    <div>
                      <p className="match-details-contact-label">Telefone</p>
                      <p className="match-details-contact-value">{match.matched_inmate.responsible.phone}</p>
                    </div>
                  </div>
                )}
              </div>
            )}

            {/* Regime */}
            <div className="match-details-custody-card">
              <div>
                <p className="match-details-custody-label">Regime de Custódia</p>
                <p className="match-details-custody-value">{getCustodyLabel(match.matched_inmate.custody)}</p>
              </div>
              <span className={`match-details-custody-badge ${getCustodyClass(match.matched_inmate.custody)}`}>
                {match.matched_inmate.custody}
              </span>
            </div>
          </div>
        </div>

        {/* Análise de Compatibilidade */}
        <div className="match-details-analysis">
          <div className="match-details-analysis-header">
            <div className="match-details-analysis-icon">
              📊
            </div>
            <h2 className="match-details-analysis-title">Análise de Compatibilidade</h2>
          </div>

          <div className="match-details-analysis-grid">
            <div className="match-details-analysis-item">
              <div className="match-details-analysis-item-icon green">
                ✓
              </div>
              <h3 className="match-details-analysis-item-title">Troca Perfeita</h3>
              <p className="match-details-analysis-item-text">
                As localidades de origem e destino são perfeitamente compatíveis para uma troca
              </p>
            </div>

            <div className="match-details-analysis-item">
              <div className="match-details-analysis-item-icon blue">
                ✓
              </div>
              <h3 className="match-details-analysis-item-title">Mesmo Regime</h3>
              <p className="match-details-analysis-item-text">
                Ambos os reeducandos estão no regime {getCustodyLabel(match.custody)}
              </p>
            </div>

            <div className="match-details-analysis-item">
              <div className="match-details-analysis-item-icon purple">
                {match.match_score}%
              </div>
              <h3 className="match-details-analysis-item-title">Score de Match</h3>
              <p className="match-details-analysis-item-text">
                Compatibilidade máxima entre os reeducandos
              </p>
            </div>
          </div>

          <div className="match-details-analysis-info">
            <div className="match-details-analysis-info-content">
              <span className="match-details-analysis-info-icon">ℹ</span>
              <div>
                <h4 className="match-details-analysis-info-title">Por que este é um match perfeito?</h4>
                <p className="match-details-analysis-info-text">
                  Este match é considerado perfeito porque ambos os reeducandos desejam realizar uma troca de localização
                  que beneficia mutuamente. O reeducando <strong>#{match.my_inmate.id}</strong> quer ir de <strong>{match.my_inmate.origin?.name}</strong> para <strong>{match.my_inmate.destination?.name}</strong>,
                  enquanto o reeducando <strong>#{match.matched_inmate.id}</strong> quer fazer exatamente o caminho inverso: de <strong>{match.matched_inmate.origin?.name}</strong> para <strong>{match.matched_inmate.destination?.name}</strong>.
                  Além disso, ambos estão no mesmo regime de custódia, facilitando a aprovação da transferência.
                </p>
              </div>
            </div>
          </div>
        </div>

        {/* Ações */}
        <div className="match-details-actions">
          <h2 className="match-details-actions-title">Próximos Passos</h2>
          <p className="match-details-actions-text">
            Entre em contato com o responsável pelo outro reeducando para iniciar o processo de troca
          </p>
          <div className="match-details-actions-buttons">
            <button
              onClick={() => {
                if (match.matched_inmate.responsible.phone) {
                  window.open(`https://wa.me/55${match.matched_inmate.responsible.phone.replace(/\D/g, '')}`, '_blank');
                }
              }}
              className="match-details-btn-whatsapp"
            >
              💬 Contatar via WhatsApp
            </button>
            <Button
              variant="secondary"
              onClick={() => navigate('/matches')}
            >
              Ver Outros Matches
            </Button>
          </div>
        </div>
      </div>
    </Layout>
  );
}
