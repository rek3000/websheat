export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      <body>
      <div>TOP BAR</div>
      {children}
      <div>BOTTOM BAR</div>
      </body>
    </html>
  )
}
