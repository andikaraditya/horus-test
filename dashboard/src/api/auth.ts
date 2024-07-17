"use server"

import fetchBackend from "./fetch"
import { cookies } from "next/headers"
type LoginResponse = {
  token: string
}

export async function fetchLogin({ username, password }: { username: string; password: string }) {
  try {
    const cookieStore = cookies()
    const result: LoginResponse = await fetchBackend({
      path: "/login",
      method: "POST",
      data: {
        username,
        password,
      },
    })
    const threeDay = 3 * 24 * 60 * 60

    cookieStore.set("token", result.token, { maxAge: threeDay })
  } catch (error: any) {
    throw new Error(error.error)
  }
}

export async function fetchRegister({
  username,
  password,
  nama,
  email,
}: {
  username: string
  password: string
  nama: string
  email: string
}) {
  try {
    const result: LoginResponse = await fetchBackend({
      path: "/register",
      method: "POST",
      data: {
        username,
        password,
        nama,
        email,
      },
    })
  } catch (error) {
    throw error
  }
}
