import type { Metadata } from "next";
import { ReactNode } from "react";

export const metadata: Metadata = {
  title: "Admin Panel | Portfolio",
  description: "Manage and edit portfolio details from this admin panel.",
};

export default function AdminLayout({ children }: { children: ReactNode }) {
  return (
    <div>
      {children}
    </div>
  );
}
