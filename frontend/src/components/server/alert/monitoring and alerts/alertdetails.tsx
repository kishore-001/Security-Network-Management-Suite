import React from 'react';
import './alertdetails.css';

interface AlertDetailsProps {
  alert: {
    id: string;
    severity: string;
    title: string;
    description: string;
    time: string;
    status: string;
    assignedTo: string;
  };
  onClose: () => void;
}

const AlertDetails: React.FC<AlertDetailsProps> = ({ alert, onClose }) => {
  return (
    <div className="modal-overlay">
      <div className="alert-modal">
        <div className="modal-header">
          <h3>🔺 Alert Details</h3>
          <button className="close-btn" onClick={onClose}>×</button>
        </div>
        <p className="modal-subtext">Detailed information about this alert</p>

        <div className="alert-title-row">
          <span className="badge red">{alert.severity}</span>
          <h4>{alert.title}</h4>
        </div>

        <div className="description-box">
          {alert.description}
        </div>

        <div className="alert-info-grid">
          <div>
            <label>Alert ID</label>
            <p>{alert.id}</p>
          </div>
          <div>
            <label>Detected</label>
            <p>{alert.time}</p>
          </div>
          <div>
            <label>Status</label>
            <p><span className="badge red">{alert.status}</span></p>
          </div>
          <div>
            <label>Assigned To</label>
            <p>{alert.assignedTo}</p>
          </div>
        </div>

        <div className="modal-actions">
          <button className="acknowledge-btn" disabled>Acknowledge</button>
          <button className="resolve-btn">Resolve Alert</button>
        </div>
      </div>
    </div>
  );
};

export default AlertDetails;
