/* eslint-disable @typescript-eslint/no-unused-vars */
"use client";
import { useState } from "react";
import { useThemeStore } from "@/app/store/themeStore";
import { Sun, Moon } from "lucide-react";
import { motion, AnimatePresence } from "framer-motion";

export default function ThemeToggle() {
  const { theme, toggleTheme } = useThemeStore();
  const [animating, setAnimating] = useState(false);

  const handleToggle = () => {
    setAnimating(true);
    setTimeout(() => {
      toggleTheme();
      setAnimating(false);
    }, 400);
  };

  return (
    <button
      onClick={handleToggle}
      className="relative p-2 m-4 border rounded bg-primary text-primary-foreground flex items-center justify-center overflow-hidden"
    >
      {/* Background overlay for smooth engulf effect */}
      <motion.div
        key={theme}
        initial={{ scale: 0, opacity: 0 }}
        animate={{ scale: 1.2, opacity: 1 }}
        exit={{ scale: 2, opacity: 0 }}
        transition={{ duration: 0.4, ease: "easeInOut" }}
        className="absolute inset-0 bg-background transition-colors duration-500 rounded-full"
      />

      {/* Icon Animation */}
      <AnimatePresence mode="wait">
        {theme === "light" ? (
          <motion.div
            key="moon"
            initial={{ scale: 0.8, opacity: 0 }}
            animate={{ scale: 1, opacity: 1 }}
            exit={{ scale: 1.2, opacity: 0 }}
            transition={{ duration: 0.3 }}
          >
            <Moon size={24} className="relative z-10" />
          </motion.div>
        ) : (
          <motion.div
            key="sun"
            initial={{ scale: 0.8, opacity: 0 }}
            animate={{ scale: 1, opacity: 1 }}
            exit={{ scale: 1.2, opacity: 0 }}
            transition={{ duration: 0.3 }}
          >
            <Sun size={24} className="relative z-10" />
          </motion.div>
        )}
      </AnimatePresence>
    </button>
  );
}
