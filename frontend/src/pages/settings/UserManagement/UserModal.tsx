import React, { useState } from 'react';
import './UserModal.css'; // move modal styles here

interface Props {
  onSubmit: (user: { name: string; email: string; role: string }) => void;
  onClose: () => void;
}

const UserModal: React.FC<Props> = ({ onClose, onSubmit }) => {
  const [name, setName] = useState('');
  const [role, setRole] = useState('');
  const [email, setEmail] = useState('');

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (!name || !role || !email) return;
    onSubmit({ name, role, email });
    onClose();
  };

  return (
    <div className="modal-overlay">
      <div className="modal-box">
        <h2>Add New User</h2>
        <form onSubmit={handleSubmit}>
          <label>
            Name
            <input
              type="text"
              placeholder="Enter name"
              value={name}
              onChange={(e) => setName(e.target.value)}
            />
          </label>
          <label>
            Email
            <input
              type="email"
              placeholder="Enter email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
            />
          </label>
          <label>
            Role
            <select value={role} onChange={(e) => setRole(e.target.value)} required>
              <option value="">Select Role</option>
              <option value="Admin">Admin</option>
              <option value="Moderator">Moderator</option>
              <option value="Viewer">Viewer</option>
              <option value="Operator">Operator</option>
            </select>
          </label>
          <div className="modal-actions">
            <button type="submit" className="btn add">Add</button>
            <button type="button" className="btn remove" onClick={onClose}>Cancel</button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default UserModal;
