import { NextResponse, NextRequest } from "next/server"
import { cookies } from "next/headers"

export function middleware(request: NextRequest) {
  const token = cookies().has("token")

  if (request.nextUrl.pathname.startsWith("/_next")) {
    return NextResponse.next()
  }

  if (!token) {
    if (
      request.nextUrl.pathname.startsWith("/login") ||
      request.nextUrl.pathname.startsWith("/register")
    ) {
      return NextResponse.next()
    } else {
      return NextResponse.redirect(new URL("/login", request.url))
    }
  }
  return NextResponse.next()
}
