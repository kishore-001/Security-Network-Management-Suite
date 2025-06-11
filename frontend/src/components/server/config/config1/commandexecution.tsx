import React, { useState } from 'react';
import './commandexecution.css';
import { FaTerminal, FaPlay } from 'react-icons/fa';
import ModalWrapper from './modalwrapper';

const CommandExecution: React.FC = () => {
  const [showModal, setShowModal] = useState(false);
  const [command, setCommand] = useState('');

  const handleExecute = () => {
    // TODO: Connect to backend for actual command execution
    console.log("Command to execute:", command);
    setShowModal(false);
    setCommand('');
  };

  return (
    <>
      <div className="card" onClick={() => setShowModal(true)}>
        <div className="card-icon green"><FaTerminal /></div>
        <h2>Command Execution</h2>
        <p>Execute administrative commands remotely</p>
      </div>

      {showModal && (
        <ModalWrapper title="Remote Command Execution" onClose={() => setShowModal(false)}>
          <p className="subtitle">Execute administrative commands remotely with elevated privileges</p>
          <label>Command Input</label>
          <textarea
            placeholder="Enter administrative command..."
            value={command}
            onChange={(e) => setCommand(e.target.value)}
          />
          <button onClick={handleExecute}><FaPlay /> Execute Command</button>
        </ModalWrapper>
      )}
    </>
  );
};

export default CommandExecution;
