import React, { useState } from 'react';
import UserTable from './Usertable';
import UserModal from './UserModal';
import './UserManagement.css';

interface User {
  id: number;
  name: string;
  email: string;
  role: string;
}

const UserManagement: React.FC = () => {
  const [users, setUsers] = useState<User[]>([]);
  const [showModal, setShowModal] = useState(false);

  const handleAddUser = (newUser: Omit<User, 'id'>) => {
    const userWithId: User = { ...newUser, id: Date.now() };
    setUsers((prev) => [...prev, userWithId]);
    setShowModal(false);
  };

  const handleRemoveUser = () => {
    if (users.length === 0) return;
    setUsers((prev) => prev.slice(0, -1)); // remove last user
  };

  return (
    <div>
      <div className="action-buttons">
        <button className="btn add" onClick={() => setShowModal(true)}>+ Add User</button>
        <button className="btn remove" onClick={handleRemoveUser}>Remove User</button>
      </div>
      <UserTable users={users} />
      {showModal && <UserModal onClose={() => setShowModal(false)} onSubmit={handleAddUser} />}
    </div>
  );
};

export default UserManagement;
