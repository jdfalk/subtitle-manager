import { useEffect, useState } from "react";
import Settings from "./Settings.jsx";
import Dashboard from "./Dashboard.jsx";
import "./App.css";

function App() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [status, setStatus] = useState("");
  const [authed, setAuthed] = useState(false);
  const [page, setPage] = useState("dashboard");

  useEffect(() => {
    fetch("/api/config").then((res) => {
      if (res.ok) setAuthed(true);
    });
  }, []);

  const login = async () => {
    const form = new URLSearchParams({ username, password });
    const res = await fetch("/api/login", { method: "POST", body: form });
    if (res.ok) {
      setStatus("logged in");
      setAuthed(true);
    } else {
      setStatus("login failed");
    }
  };

  if (!authed) {
    return (
      <div className="login">
        <h1>Subtitle Manager</h1>
        <input
          placeholder="Username"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
        />
        <input
          type="password"
          placeholder="Password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
        <button onClick={login}>Login</button>
        <p>{status}</p>
      </div>
    );
  }

  return (
    <div>
      <nav>
        <button onClick={() => setPage("dashboard")}>Dashboard</button>
        <button onClick={() => setPage("settings")}>Settings</button>
      </nav>
      {page === "settings" ? <Settings /> : <Dashboard />}
    </div>
  );
}

export default App;
