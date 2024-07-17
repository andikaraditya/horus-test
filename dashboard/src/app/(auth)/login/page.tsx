"use client"

import { fetchLogin } from "@/api/auth"
import { useRouter } from "next/navigation"
import { ChangeEvent, FormEvent, useState } from "react"
import Link from "next/link"

function LoginPage() {
  const router = useRouter()
  const [loginForm, setLoginForm] = useState({
    username: "",
    password: "",
  })
  const [error, setError] = useState("")
  function handleFormChange(e: ChangeEvent<HTMLInputElement>) {
    const { name, value } = e.target
    setLoginForm((prevValue) => ({
      ...prevValue,
      [name]: value,
    }))
  }
  async function handleLogin(e: FormEvent) {
    e.preventDefault()
    try {
      await fetchLogin(loginForm)
      router.push("/home")
    } catch (error: any) {
      setError(error.message)
    }
  }
  return (
    <div className="h-screen flex items-center justify-center text-slate-800">
      <div className="w-2/3 h-2/3 lg:w-[400px] lg:h-[500px] border-[0.2rem] bg-slate-50 border-slate-800 rounded-xl shadow-lg">
        <h1 className="text-center text-4xl font-bold mt-10">Login</h1>
        <p className="text-center mt-4">
          click{" "}
          <Link className="underline" href="/register">
            here
          </Link>{" "}
          to register
        </p>
        <p className="text-red-600 font-semibold text-center mt-2">{error}</p>
        <form onSubmit={handleLogin} className="py-5 px-7">
          <div className="my-8">
            <label className="block" htmlFor="username">
              Username
            </label>
            <input
              onChange={handleFormChange}
              className="form-input"
              type="text"
              name="username"
              id="username"
              required
            />
          </div>
          <div className="my-8">
            <label className="block" htmlFor="password">
              Password
            </label>
            <input
              onChange={handleFormChange}
              className="form-input"
              type="password"
              name="password"
              id="password"
              required
            />
          </div>
          <div className="my-8 flex justify-center">
            <button className="bg-white border-[0.1rem] border-slate-300 px-8 py-2 rounded-md mx-auto">
              Login
            </button>
          </div>
        </form>
      </div>
    </div>
  )
}

export default LoginPage
