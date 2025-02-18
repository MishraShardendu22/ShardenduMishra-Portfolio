"use client";

import { useThemeStore } from "@/app/store/themeStore";

export default function ThemeToggle() {
  const { theme, toggleTheme } = useThemeStore();

  return (
    <button
      onClick={toggleTheme}
      className="p-2 m-4 border rounded bg-primary text-primary-foreground"
    >
      {theme === "light" ? "Dark Mode" : "Light Mode"}
    </button>
  );
}
