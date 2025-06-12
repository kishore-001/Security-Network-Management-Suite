import React, { useState } from 'react';
import './active alerts.css';
import AlertDetails from './alertdetails';

const ActiveAlerts: React.FC = () => {
  const [selectedAlert, setSelectedAlert] = useState<any | null>(null);

  const handleDetailsClick = () => {
    setSelectedAlert({
      id: 'ALT-000001',
      severity: 'high',
      title: 'CPU usage exceeded 90% on server-01',
      description: 'CPU usage exceeded 90% on server-01 for 10 minutes',
      time: '10 minutes ago',
      status: 'Active',
      assignedTo: 'Unassigned',
    });
  };

  const closeModal = () => {
    setSelectedAlert(null);
  };

  return (
    <>
      <div className="alerts-table-card">
        <div className="table-header">
          <h3>🔺 Active Alerts</h3>
          <button className="filter-btn">All Alerts ⌄</button>
        </div>
        <table className="alerts-table">
          <thead>
            <tr>
              <th>Severity</th>
              <th>Alert</th>
              <th>Time</th>
              <th>Actions</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td><span className="badge red">high</span></td>
              <td>⚙️ CPU usage exceeded 90% on server-01</td>
              <td>🕒 10 minutes ago</td>
              <td>
                <button className="details-btn" onClick={handleDetailsClick}>
                  👁️ Details
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      {selectedAlert && <AlertDetails alert={selectedAlert} onClose={closeModal} />}
    </>
  );
};

export default ActiveAlerts;
