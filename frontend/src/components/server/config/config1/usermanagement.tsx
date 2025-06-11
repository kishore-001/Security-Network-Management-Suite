import React, { useState } from 'react';
import './usermanagement.css';
import ModalWrapper from './modalwrapper';
import { FaUserPlus, FaTrash } from 'react-icons/fa';

interface User {
  id: string;
  username: string;
  created: string;
  status: 'Active' | 'Inactive';
}

const UserManagement: React.FC = () => {
  const [showModal, setShowModal] = useState(false);
  const [users, setUsers] = useState<User[]>([
    { id: 'SRV-001', username: 'admin', created: '2024-01-15', status: 'Active' },
    { id: 'SRV-002', username: 'developer', created: '2024-01-20', status: 'Active' },
    { id: 'SRV-003', username: 'operator', created: '2024-02-01', status: 'Inactive' },
  ]);
  const [newUsername, setNewUsername] = useState('');

  const addUser = () => {
    if (!newUsername.trim()) return;
    const newUser: User = {
      id: `SRV-00${users.length + 1}`,
      username: newUsername.trim(),
      created: new Date().toISOString().split('T')[0],
      status: 'Active',
    };
    setUsers([...users, newUser]);
    setNewUsername('');
  };

  const deleteUser = (id: string) => {
    setUsers(users.filter((user) => user.id !== id));
  };

  return (
    <>
      <div className="card" onClick={() => setShowModal(true)}>
        <div className="card-icon purple"><FaUserPlus /></div>
        <h2>User Management</h2>
        <p>Manage server users, permissions, and access control</p>
      </div>

      {showModal && (
        <ModalWrapper title="User Management" onClose={() => setShowModal(false)}>
          <p className="subtitle">Manage server users, permissions, and access control</p>

          <div className="user-input-row">
            <input
              type="text"
              placeholder="Enter username"
              value={newUsername}
              onChange={(e) => setNewUsername(e.target.value)}
            />
            <button onClick={addUser}>
              <FaUserPlus /> Add User
            </button>
          </div>

          <table>
            <thead>
              <tr>
                <th>Server ID</th>
                <th>Username</th>
                <th>Created</th>
                <th>Status</th>
                <th>Actions</th>
              </tr>
            </thead>
            <tbody>
              {users.map((user) => (
                <tr key={user.id}>
                  <td>{user.id}</td>
                  <td>{user.username}</td>
                  <td>{user.created}</td>
                  <td>
                    <span className={`status ${user.status.toLowerCase()}`}>
                      {user.status}
                    </span>
                  </td>
                  <td>
                    <button className="delete-btn" onClick={() => deleteUser(user.id)}>
                      <FaTrash />
                    </button>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </ModalWrapper>
      )}
    </>
  );
};

export default UserManagement;
