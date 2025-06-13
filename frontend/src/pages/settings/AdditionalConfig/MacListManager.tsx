import React, { useState } from 'react';
import MacModal from './MacModal';
import type { MacEntry } from './MacModal';


import './AdditionalConfig.css';

const MacListManager: React.FC = () => {
  const [showModal, setShowModal] = useState(false);
  const [macList, setMacList] = useState<MacEntry[]>([]);

  const handleAddMac = (entry: MacEntry) => {
    setMacList([...macList, entry]);
  };

  return (
    <div className="mac-list-container">
      <div className="action-buttons">
        <button className="btn add" onClick={() => setShowModal(true)}>+ Add MAC Address</button>
        <button className="btn remove">Remove MAC Address</button>
      </div>

      <div className="dual-table-container">
        <div className="mac-table">
          <h4>Whitelist ({macList.filter(e => e.type === 'Whitelist').length})</h4>
          {macList.filter(e => e.type === 'Whitelist').length === 0 ? (
            <div className="empty-table">
              <p>No MAC addresses found</p>
              <button className="btn add" onClick={() => setShowModal(true)}>Add MAC Address</button>
            </div>
          ) : (
            <ul>
              {macList.filter(e => e.type === 'Whitelist').map((entry, i) => (
                <li key={i}>{entry.mac} {entry.note && `- ${entry.note}`}</li>
              ))}
            </ul>
          )}
        </div>

        <div className="mac-table">
          <h4>Blacklist ({macList.filter(e => e.type === 'Blacklist').length})</h4>
          {macList.filter(e => e.type === 'Blacklist').length === 0 ? (
            <div className="empty-table">
              <p>No MAC addresses found</p>
              <button className="btn add" onClick={() => setShowModal(true)}>Add MAC Address</button>
            </div>
          ) : (
            <ul>
              {macList.filter(e => e.type === 'Blacklist').map((entry, i) => (
                <li key={i}>{entry.mac} {entry.note && `- ${entry.note}`}</li>
              ))}
            </ul>
          )}
        </div>
      </div>

      {showModal && (
        <MacModal
          onClose={() => setShowModal(false)}
          onSave={handleAddMac}
        />
      )}
    </div>
  );
};

export default MacListManager;
