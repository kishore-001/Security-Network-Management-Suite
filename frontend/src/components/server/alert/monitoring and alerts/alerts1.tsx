import React from 'react';
import './alerts1.css';
import Alerts from './alerts';
import Issues from './issues';
import ActiveAlerts from './active alerts';

const Alerts1: React.FC = () => {
  return (
    <div className="monitoring-dashboard">
      <h1>Monitoring Dashboard</h1>
      <p className="subtitle">Real-time system monitoring and alert management</p>

      <div className="card-row">
        <Issues />
        <Alerts />
      </div>

      {/* Active Alerts Table with modal logic inside */}
      <ActiveAlerts />
    </div>
  );
};

export default Alerts1;
