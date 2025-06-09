import './commandexecution.css';
const CommandExecution = () => {
  return (
    <div className="card">
      <h2>Command Execution</h2>
      <textarea placeholder="Enter shell command..." />
      <button>Run Command</button>
    </div>
  );
};

export default CommandExecution;
