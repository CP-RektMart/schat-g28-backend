'use client'

import { MessageCircle } from 'lucide-react'
import { redirect } from 'next/navigation'

import { Button } from '@/components/ui/button'

export default function LoginPage() {
  // This would be replaced with actual auth logic
  const handleLogin = () => {
    // Simulate login and redirect
    redirect('/chat')
  }

  return (
    <div className='flex min-h-screen flex-col items-center justify-center bg-gray-50'>
      <div className='w-full max-w-md space-y-8 rounded-lg bg-white p-8 shadow-md'>
        <div className='text-center'>
          <div className='mx-auto flex h-16 w-16 items-center justify-center rounded-full bg-gray-100'>
            <MessageCircle className='h-8 w-8 text-gray-600' />
          </div>
          <h2 className='mt-6 text-3xl font-bold text-gray-900'>ChatApp</h2>
          <p className='mt-2 text-sm text-gray-600'>
            Connect with friends and colleagues
          </p>
        </div>

        <div className='mt-8 space-y-4'>
          <Button
            onClick={handleLogin}
            className='flex w-full items-center justify-center gap-2 bg-white text-gray-700 hover:bg-gray-50'
            variant='outline'
          >
            <span>Sign in with Google</span>
          </Button>

          <div className='text-center text-xs text-gray-500'>
            By signing in, you agree to our Terms of Service and Privacy Policy
          </div>
        </div>
      </div>
    </div>
  )
}
