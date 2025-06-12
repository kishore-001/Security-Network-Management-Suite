import React from 'react';
import './alertdetails.css';

interface AlertDetailsProps {
  alert: any;
  onClose: () => void;
}

const AlertDetails: React.FC<AlertDetailsProps> = ({ alert, onClose }) => {
  if (!alert) return null;

  return (
    <div className="modal-overlay">
      <div className="alert-modal">
        <div className="modal-header">
          <div className="modal-header-content">
            <div className="modal-title-row">
              <span className="modal-icon">⚠️</span>
              <strong>Alert Details</strong>
            </div>
            <p className="subtext">Detailed information about this alert</p>
          </div><div>
          <button className="close-btn" onClick={onClose}>×</button></div>
        </div>

        <div className="alert-title-row">
          <span className="badge red">high</span>
          <h4>CPU usage exceeded 90% on server-01</h4>
        </div>

        <div className="description-box">
          CPU usage exceeded 90% on server-01 for 10 minutes
        </div>

        <div className="alert-info-grid">
          <div>
            <label>Alert ID</label>
            <p>ALT-000001</p>
          </div>
          <div>
            <label>Detected</label>
            <p>10 minutes ago</p>
          </div>
          <div>
            <label>Status</label>
            <p><span className="badge red">Active</span></p>
          </div>
          <div>
            <label>Assigned To</label>
            <p>Unassigned</p>
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
