// file: webui/src/__tests__/App.test.jsx
import "@testing-library/jest-dom/vitest";
import { fireEvent, render, screen, waitFor } from "@testing-library/react";
import { beforeEach, describe, expect, test, vi } from "vitest";
import App from "../App.jsx";

describe("App component", () => {
  beforeEach(() => {
    vi.restoreAllMocks();
    global.fetch = vi.fn(() =>
      Promise.resolve({
        ok: true,
        json: () =>
          Promise.resolve({ running: false, completed: 0, files: [] }),
      }),
    );
  });

  test("shows login form when unauthenticated", async () => {
    fetch.mockResolvedValueOnce({ ok: false });
    fetch.mockResolvedValueOnce({
      json: () => Promise.resolve({ needed: false }),
    });
    render(<App />);
    expect(screen.getByText("Subtitle Manager")).toBeInTheDocument();
  });

  test("successful login renders dashboard", async () => {
    fetch.mockResolvedValueOnce({ ok: false }); // config check
    fetch.mockResolvedValueOnce({
      json: () => Promise.resolve({ needed: false }),
    });
    render(<App />);
    fetch.mockResolvedValueOnce({ ok: true });
    fireEvent.click(screen.getAllByText("Login")[0]);
    await waitFor(() =>
      expect(fetch).toHaveBeenLastCalledWith("/api/login", expect.any(Object)),
    );
  });
});
