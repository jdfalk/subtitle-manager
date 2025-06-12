// file: webui/src/__tests__/Settings.test.jsx
import "@testing-library/jest-dom/vitest";
import { fireEvent, render, screen, waitFor } from "@testing-library/react";
import { beforeEach, describe, expect, test, vi } from "vitest";
import Settings from "../Settings.jsx";

describe("Settings component", () => {
  beforeEach(() => {
    vi.restoreAllMocks();
    global.fetch = vi.fn(() => Promise.resolve({ ok: true, json: () => Promise.resolve({}) }));
  });

  test("loads and saves configuration", async () => {
    fetch.mockResolvedValueOnce({
      ok: true,
      json: () => Promise.resolve({ foo: "bar" }),
    });
    render(<Settings />);
    await screen.findByDisplayValue("bar");
    fetch.mockResolvedValueOnce({ ok: true });
    fireEvent.change(screen.getByDisplayValue("bar"), {
      target: { value: "baz" },
    });
    fireEvent.click(screen.getByText("Save"));
    await waitFor(() =>
      expect(fetch).toHaveBeenLastCalledWith("/api/config", expect.any(Object)),
    );
  });
});
