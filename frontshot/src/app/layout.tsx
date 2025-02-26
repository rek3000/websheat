import type { Metadata } from "next";

import localFont from "next/font/local";
import "./globals.css";
import { Navbar } from "@/components/Navbar/Navbar"
import { ThemeProvider } from "@/components/ThemeProvider/ThemeProvider";

const systemFont = localFont({ src: "../assets/DepartureMono-Regular.otf", display: "swap" });

export const metadata: Metadata = {
  title: "Websheat",
  description: "dogshit web::front::shot",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html className={`bg-gray-500 lg:min-h-60 md:min-h-screen max-w-fit m-auto justify-center`} lang="en" suppressHydrationWarning>
      <body
        className={`${systemFont.className} flex items-center flex-col border-4`}>
        <ThemeProvider attribute="class" defaultTheme="system" enableSystem disableTransitionOnChange>
          <div className={`m-2`}>
            <Navbar />
            {children}
          </div>
        </ThemeProvider>
      </body>
    </html>
  );
}
