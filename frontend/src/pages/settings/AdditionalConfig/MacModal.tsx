import React, { useState } from 'react';
import './AdditionalConfig.css';

export interface MacEntry {
  mac: string;
  type: 'Whitelist' | 'Blacklist';
  note?: string;
}

interface MacModalProps {
  onClose: () => void;
  onSave: (entry: MacEntry) => void;
}

const MacModal: React.FC<MacModalProps> = ({ onClose, onSave }) => {
  const [mac, setMac] = useState('');
  const [type, setType] = useState<'Whitelist' | 'Blacklist'>('Whitelist');
  const [note, setNote] = useState('');

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();

    const macRegex = /^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$/;
    if (!macRegex.test(mac)) {
      alert('Invalid MAC address format');
      return;
    }

    onSave({ mac, type, note });
    onClose();
  };

  return (
    <div className="modal-overlay">
      <form className="modal-content" onSubmit={handleSubmit}>
        <h3>Add MAC Address</h3>

        <input
          type="text"
          placeholder="XX:XX:XX:XX:XX:XX"
          value={mac}
          onChange={(e) => setMac(e.target.value)}
          required
        />

        <select value={type} onChange={(e) => setType(e.target.value as 'Whitelist' | 'Blacklist')}>
          <option value="Whitelist">Whitelist</option>
          <option value="Blacklist">Blacklist</option>
        </select>

        <input
          type="text"
          placeholder="Note (optional)"
          value={note}
          onChange={(e) => setNote(e.target.value)}
        />

        <div className="modal-actions">
          <button type="submit" className="btn add">Save</button>
          <button type="button" className="btn remove" onClick={onClose}>Cancel</button>
        </div>
      </form>
    </div>
  );
};

export default MacModal;
