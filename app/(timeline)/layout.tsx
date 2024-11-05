'use server'

import { Navigation } from '@/components/navigation'
import { TweetDialog } from '@/components/tweet'
import { auth } from '@/lib/auth'
import { UserProvider } from '@/providers/user'
import { Avatar, Button, Text, TextField } from '@radix-ui/themes'
import { Search } from 'lucide-react'
import Image from 'next/image'
import Link from 'next/link'

const Sidebar = ({ user }) => {
  return (
    <div className='w-64 pr-6 space-y-6 sticky top-0 h-screen'>
      <div className='flex items-center space-x-3 mb-8 pl-2 pt-5'>
        <Avatar
          src={user?.avatar!}
          alt={user?.name!}
          className='w-12 h-12'
          fallback={user?.username?.charAt(0)!}
        />
      </div>

      <Navigation user={user} />

      <TweetDialog />
    </div>
  )
}

export default async function Page({
  children,
}: {
  children: React.ReactNode
}) {
  const user = await auth()

  return (
    <UserProvider user={user}>
      <div className='flex justify-center items-center min-h-screen'>
        <div className='flex item w-full max-w-max mx-auto'>
          {/* Left Sidebar */}
          {user && <Sidebar user={user} />}
          {!user && (
            <div className='w-64 mt-10 pr-6 space-y-2 sticky top-0 h-screen'>
              <Text size='3'>サインアップまたはログインして会話に参加</Text>
              <Button asChild className='w-full h-10 font-bold'>
                <Link href='/auth'>ログイン</Link>
              </Button>
            </div>
          )}

          {/* Main Content */}
          <div className='w-[600px] border-x'>
            <div className='border-b sticky top-0 z-10 backdrop-blur-sm'>
              <Image
                src='/icon.svg'
                alt='Logo'
                width={64}
                height={64}
                className='mx-auto p-4'
              />
            </div>

            <div className='p-4'>{children}</div>
          </div>

          {/* Right Sidebar */}
          <div className='w-64 pl-6 sticky top-0 h-screen'>
            <div className='sticky top-4'>
              <TextField.Root
                size='3'
                variant='soft'
                color='gray'
                placeholder='Search'
                className='w-full'
              >
                <TextField.Slot>
                  <Search className='w-5 h-5 text-gray-500' />
                </TextField.Slot>
              </TextField.Root>

              <div className='mt-4 bg-gray-50 rounded-2xl p-4 space-y-4'>
                <h2 className='font-semibold text-xl'>Following</h2>
                <h2 className='font-semibold text-xl'>Discover</h2>
                <h2 className='font-semibold text-xl'>For You</h2>
              </div>
            </div>
          </div>
        </div>
      </div>
    </UserProvider>
  )
}
