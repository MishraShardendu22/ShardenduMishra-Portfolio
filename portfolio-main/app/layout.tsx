import ThemeProvider from "@/components/ThemeProvider";
import ThemeToggle from "@/components/ThemeToggle";
import { Toaster } from "react-hot-toast";
import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  title: "Shardendu Mishra | Portfolio Website",
  description:
    "Hi I am Shardendu Mishra, A CSE student, a web developer and a tech Enthusiast.",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
        <html lang="en">
      <body className="antialiased">
        <ThemeProvider>
          <header className="flex justify-end">
            <ThemeToggle />
          </header>
          <main className="container mx-auto px-4">{children}</main>
          <Toaster 
            position="bottom-left" 
            reverseOrder={false} 
            gutter={8} 
            toastOptions={{
              className: '',
              style: {
                margin: '40px',
                background: '#363636',
                color: '#fff',
                fontSize: '1.5rem',
              },
            }}
          />
        </ThemeProvider>
      </body>
    </html>
  );
}
