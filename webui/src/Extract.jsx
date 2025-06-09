import { useState } from "react";

/**
 * Extract provides a simple form to request subtitle extraction for a media file.
 * The path to the media file is POSTed to `/api/extract` and the number of
 * extracted items is displayed.
 */
export default function Extract() {
  const [path, setPath] = useState("");
  const [status, setStatus] = useState("");

  const doExtract = async () => {
    const res = await fetch("/api/extract", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ path }),
    });
    if (res.ok) {
      const items = await res.json();
      setStatus(`extracted ${items.length} items`);
    } else {
      setStatus("error");
    }
  };

  return (
    <div className="extract">
      <h1>Extract Subtitles</h1>
      <input
        placeholder="/path/to/media"
        value={path}
        onChange={(e) => setPath(e.target.value)}
      />
      <button onClick={doExtract}>Extract</button>
      <p>{status}</p>
    </div>
  );
}
