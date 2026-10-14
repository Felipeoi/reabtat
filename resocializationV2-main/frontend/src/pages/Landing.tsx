import { Link } from 'react-router-dom';
import '../styles/landing.css';

const plans = [
  { code: 'basic', name: 'Básico', matches: 1, price: 'R$ 29,90' },
  { code: 'professional', name: 'Profissional', matches: 3, price: 'R$ 69,90', featured: true },
  { code: 'office', name: 'Escritório', matches: 5, price: 'R$ 99,90' },
  { code: 'premium', name: 'Premium', matches: 10, price: 'R$ 179,90' },
];

const steps = [
  {
    n: '01',
    title: 'Cadastre seus reeducandos',
    text: 'Informe unidade de origem, destinos desejados e o regime. O cadastro é gratuito e ilimitado.',
  },
  {
    n: '02',
    title: 'Receba correlações automáticas',
    text: 'O sistema cruza origem, destino e regime com casos de outros escritórios e avisa quando há match.',
  },
  {
    n: '03',
    title: 'Desbloqueie o contato',
    text: 'Com um crédito, você libera o telefone do outro advogado para articular a transferência.',
  },
];

const features = [
  {
    title: 'Match por origem e destino',
    text: 'Correlação quando a origem de um reeducando é destino do outro, e vice-versa.',
  },
  {
    title: 'Mesmo regime prisional',
    text: 'Só entram na lista casos com o mesmo regime: fechado, semiaberto ou aberto.',
  },
  {
    title: 'Entre escritórios diferentes',
    text: 'Matches acontecem apenas entre advogados distintos, para articulação real de defesa.',
  },
  {
    title: 'Cadastro com OAB e CPF',
    text: 'A plataforma é exclusiva para advogados. O cadastro exige inscrição ativa e CPF válido.',
  },
  {
    title: 'Aviso por e-mail',
    text: 'Quando surge um match novo, você recebe notificação para não perder a janela de contato.',
  },
  {
    title: 'Créditos via PIX',
    text: 'Pague só quando quiser o contato. Planos de 1 a 10 matches, com QR Code e copia e cola.',
  },
];

export default function Landing() {
  return (
    <div className="landing">
      <header className="landing-nav">
        <div className="landing-nav-inner">
          <a href="#inicio" className="landing-brand">
            <img
              src="/rehabtat-logo.png"
              alt="Rehab TaT - Reabilitação Transferência a Transferência"
              className="landing-logo"
            />
          </a>
          <nav className="landing-nav-links" aria-label="Seções">
            <a href="#como-funciona">Como funciona</a>
            <a href="#recursos">Recursos</a>
            <a href="#planos">Planos</a>
          </nav>
          <div className="landing-nav-actions">
            <Link to="/login" className="btn btn-secondary btn-sm">
              Entrar
            </Link>
            <Link to="/signup" className="btn btn-primary btn-sm">
              Criar conta
            </Link>
          </div>
        </div>
      </header>

      <main>
        <section id="inicio" className="landing-hero">
          <div className="landing-hero-inner">
            <p className="landing-kicker">Para advogados com OAB ativa</p>
            <h1>
              Transferências prisionais com o escritório certo, no momento certo.
            </h1>
            <p className="landing-lead">
              O Rehab TaT cruza origem, destino e regime dos seus reeducandos com
              casos de outros advogados. Quando há compatibilidade, você decide
              se libera o contato.
            </p>
            <div className="landing-hero-actions">
              <Link to="/signup" className="btn btn-primary btn-lg">
                Começar agora
              </Link>
              <a href="#como-funciona" className="btn btn-secondary btn-lg">
                Ver como funciona
              </a>
            </div>
            <ul className="landing-trust">
              <li>Cadastro gratuito</li>
              <li>Você só paga para desbloquear contato</li>
              <li>Pagamento via PIX</li>
            </ul>
          </div>
        </section>

        <section id="como-funciona" className="landing-section">
          <div className="landing-wrap">
            <p className="landing-kicker">Passo a passo</p>
            <h2>Três etapas até o contato</h2>
            <p className="landing-section-lead">
              O cadastro de reeducandos é livre. O crédito entra só quando
              existe match e você quer falar com a outra defesa.
            </p>
            <ol className="landing-steps">
              {steps.map((step) => (
                <li key={step.n} className="landing-step">
                  <span className="landing-step-n">{step.n}</span>
                  <h3>{step.title}</h3>
                  <p>{step.text}</p>
                </li>
              ))}
            </ol>
          </div>
        </section>

        <section id="recursos" className="landing-section landing-section-alt">
          <div className="landing-wrap">
            <p className="landing-kicker">A plataforma</p>
            <h2>Feita para articular transferências</h2>
            <p className="landing-section-lead">
              Cada correlação considera os mesmos critérios que o escritório
              já usa no dia a dia: unidade, destino desejado e regime.
            </p>
            <div className="landing-features">
              {features.map((feature) => (
                <article key={feature.title} className="landing-feature">
                  <h3>{feature.title}</h3>
                  <p>{feature.text}</p>
                </article>
              ))}
            </div>
          </div>
        </section>

        <section id="planos" className="landing-section">
          <div className="landing-wrap">
            <p className="landing-kicker">Créditos</p>
            <h2>Pague apenas pelo contato</h2>
            <p className="landing-section-lead">
              Cada crédito libera o telefone do outro advogado em um match.
              Escolha o pacote que cabe no volume do seu escritório.
            </p>
            <div className="landing-plans">
              {plans.map((plan) => (
                <article
                  key={plan.code}
                  className={`landing-plan${plan.featured ? ' featured' : ''}`}
                >
                  {plan.featured && <span className="landing-plan-badge">Mais popular</span>}
                  <h3>{plan.name}</h3>
                  <p className="landing-plan-matches">
                    {plan.matches} {plan.matches === 1 ? 'match' : 'matches'}
                  </p>
                  <p className="landing-plan-price">{plan.price}</p>
                  <Link to="/signup" className="btn btn-primary">
                    Criar conta
                  </Link>
                </article>
              ))}
            </div>
          </div>
        </section>

        <section className="landing-cta-band">
          <div className="landing-wrap landing-cta-inner">
            <h2>Pronto para correlacionar seus casos?</h2>
            <p>
              Crie sua conta com OAB e CPF, cadastre o primeiro reeducando
              e deixe o Rehab TaT procurar o match.
            </p>
            <div className="landing-hero-actions">
              <Link to="/signup" className="btn btn-primary btn-lg">
                Criar conta
              </Link>
              <Link to="/login" className="btn btn-secondary btn-lg">
                Já tenho conta
              </Link>
            </div>
          </div>
        </section>
      </main>

      <footer className="landing-footer">
        <div className="landing-wrap landing-footer-inner">
          <img
            src="/rehabtat-logo.png"
            alt=""
            className="landing-footer-logo"
          />
          <p>Rehab TaT - Reabilitação Transferência a Transferência</p>
          <p className="landing-footer-copy">
            Uso exclusivo para advogados.
          </p>
        </div>
      </footer>
    </div>
  );
}
