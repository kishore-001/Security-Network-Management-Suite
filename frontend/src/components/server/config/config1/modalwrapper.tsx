import React from 'react';
import type { ReactNode } from 'react';

import './modalwrapper.css';
import { FaTimes } from 'react-icons/fa';

interface Props {
  title: string;
  children: ReactNode;
  onClose: () => void;
}

const ModalWrapper: React.FC<Props> = ({ title, children, onClose }) => {
  return (
    <div className="modal-overlay">
      <div className="modal">
        <div className="modal-header">
          <h2>{title}</h2>
          <button className="close-btn" onClick={onClose}><FaTimes /></button>
        </div>
        <div className="modal-content">{children}</div>
      </div>
    </div>
  );
};

export default ModalWrapper;
