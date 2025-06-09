// file: webui/src/Dashboard.jsx
import { useEffect, useState } from "react";

export default function Dashboard() {
  const [status, setStatus] = useState({ running: false, completed: 0, files: [] });
  const [dir, setDir] = useState("");
  const [lang, setLang] = useState("en");
  const [provider, setProvider] = useState("generic");

  const poll = async () => {
    const res = await fetch("/api/scan/status");
    if (res.ok) {
      const data = await res.json();
      setStatus(data);
      if (data.running) {
        setTimeout(poll, 1000);
      }
    }
  };

  const start = async () => {
    const body = { provider, directory: dir, lang };
    const res = await fetch("/api/scan", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(body),
    });
    if (res.ok) poll();
  };

  useEffect(() => {
    poll();
  }, []);

  return (
    <div className="dashboard">
      <h1>Dashboard</h1>
      <div>
        <input placeholder="Directory" value={dir} onChange={(e) => setDir(e.target.value)} />
        <input placeholder="Language" value={lang} onChange={(e) => setLang(e.target.value)} />
        <input placeholder="Provider" value={provider} onChange={(e) => setProvider(e.target.value)} />
        <button onClick={start}>Scan</button>
      </div>
      <p>Running: {status.running ? "yes" : "no"} ({status.completed})</p>
      <ul>
        {status.files.map((f) => (
          <li key={f}>{f}</li>
        ))}
      </ul>
    </div>
  );
}
