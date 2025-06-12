import { useEffect, useState } from "react";

export default function System() {
  const [logs, setLogs] = useState([]);
  const [info, setInfo] = useState({});
  const [tasks, setTasks] = useState({});

  useEffect(() => {
    fetch("/api/logs").then(r => r.json()).then(setLogs);
    fetch("/api/system").then(r => r.json()).then(setInfo);
    fetch("/api/tasks").then(r => r.json()).then(setTasks);
  }, []);

  return (
    <div className="system">
      <h1>System</h1>
      <h2>Logs</h2>
      <pre data-testid="logs">{logs.join("\n")}</pre>
      <h2>Tasks</h2>
      <pre data-testid="tasks">{JSON.stringify(tasks)}</pre>
      <h2>Info</h2>
      <pre data-testid="info">{JSON.stringify(info)}</pre>
    </div>
  );
}
