"use client"
import { fetchRegister } from "@/api/auth"
import Link from "next/link"
import { useRouter } from "next/navigation"
import { ChangeEvent, FormEvent, useState } from "react"

function RegisterPage() {
  const router = useRouter()
  const [registerForm, setRegisterForm] = useState({
    username: "",
    password: "",
    nama: "",
    email: "",
  })

  function handleFormChange(e: ChangeEvent<HTMLInputElement>) {
    const { name, value } = e.target
    setRegisterForm((prevValue) => ({
      ...prevValue,
      [name]: value,
    }))
    console.log(registerForm)
  }

  async function handleRegister(e: FormEvent) {
    e.preventDefault()
    try {
      await fetchRegister(registerForm)
      router.push("/login")
    } catch (error) {
      console.log(error)
    }
  }

  return (
    <div className="h-screen flex items-center justify-center text-slate-800">
      <div className="w-2/3 h-2/3 lg:w-[400px] lg:min-h-[500px] border-[0.2rem] bg-slate-50 border-slate-800 rounded-xl shadow-lg">
        <h1 className="text-center text-4xl font-bold my-7">Register Account</h1>
        <p className="text-center mt-4">
          click{" "}
          <Link className="underline" href="/login">
            here
          </Link>{" "}
          to login
        </p>
        <form onSubmit={handleRegister} className="py-5 px-7">
          <div className="my-8">
            <label className="block" htmlFor="username">
              Username
            </label>
            <input
              onChange={handleFormChange}
              required
              className="form-input"
              type="text"
              name="username"
              id="username"
            />
          </div>
          <div className="my-8">
            <label className="block" htmlFor="password">
              Password
            </label>
            <input
              onChange={handleFormChange}
              required
              className="form-input"
              type="password"
              name="password"
              id="password"
            />
          </div>
          <div className="my-8">
            <label className="block" htmlFor="nama">
              Nama
            </label>
            <input
              onChange={handleFormChange}
              className="form-input"
              type="text"
              name="nama"
              id="nama"
            />
          </div>
          <div className="my-8">
            <label className="block" htmlFor="email">
              Email
            </label>
            <input
              onChange={handleFormChange}
              className="form-input"
              type="email"
              name="email"
              id="email"
            />
          </div>
          <div className="my-8 flex justify-center">
            <button className="bg-white border-[0.1rem] border-slate-300 px-8 py-2 rounded-md mx-auto">
              Register
            </button>
          </div>
        </form>
      </div>
    </div>
  )
}

export default RegisterPage
