import { useEffect, useState } from "react";

/**
 * Wanted provides an interface for searching subtitles and maintaining
 * a list of wanted items. Results are fetched from `/api/search` and
 * selections are POSTed to `/api/wanted`.
 */
export default function Wanted() {
  const [provider, setProvider] = useState("generic");
  const [path, setPath] = useState("");
  const [lang, setLang] = useState("en");
  const [results, setResults] = useState([]);
  const [wanted, setWanted] = useState([]);

  useEffect(() => {
    fetch("/api/wanted")
      .then((r) => (r.ok ? r.json() : []))
      .then(setWanted);
  }, []);

  const search = async () => {
    const params = new URLSearchParams({ provider, path, lang });
    const res = await fetch(`/api/search?${params.toString()}`);
    if (res.ok) {
      setResults(await res.json());
    }
  };

  const add = async (url) => {
    const res = await fetch("/api/wanted", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ url }),
    });
    if (res.ok) {
      setWanted([...wanted, url]);
    }
  };

  return (
    <div className="wanted">
      <h1>Wanted</h1>
      <div>
        <input
          placeholder="Provider"
          value={provider}
          onChange={(e) => setProvider(e.target.value)}
        />
        <input
          placeholder="Media path"
          value={path}
          onChange={(e) => setPath(e.target.value)}
        />
        <input
          placeholder="Language"
          value={lang}
          onChange={(e) => setLang(e.target.value)}
        />
        <button onClick={search}>Search</button>
      </div>
      <ul>
        {results.map((u) => (
          <li key={u}>
            {u} <button onClick={() => add(u)}>Add</button>
          </li>
        ))}
      </ul>
      <h2>Wanted List</h2>
      <ul>
        {wanted.map((u) => (
          <li key={u}>{u}</li>
        ))}
      </ul>
    </div>
  );
}
