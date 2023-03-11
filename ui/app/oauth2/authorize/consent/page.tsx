'use client'
import { useSearchParams } from 'next/navigation'

import { AuthorizeRequestSchema } from '../../../../lib/oauth2/types'
import { buildAuthorizePath } from '../../../../pages/api/oauth2/route'

const AuthorizePage = () => {
  const searchParams = useSearchParams()
  if (!searchParams) {
    throw new Error('unexpected error: searchParams is null')
  }

  const parsed = AuthorizeRequestSchema.safeParse(Object.fromEntries(searchParams.entries()))
  console.log(parsed)
  if (!parsed.success) {
    // TODO: Error Handling
    console.warn(parsed.error)
    // throw parsed.error
    return <div />
  }
  const parameter = parsed.data

  return (
    <div className="m-8">
      <h2 className="mb-4 text-xl">Consent Page</h2>
      <form action={buildAuthorizePath()} method="post" className="w-full max-w-sm">
        <div className="mb-4 md:flex md:items-center">
          <Label htmlFor="client_id" />
          <Input name="client_id" value={parameter.client_id} />
        </div>
        <div className="mb-4 md:flex md:items-center">
          <Label htmlFor="redirect_uri" />
          <Input name="redirect_uri" value={parameter.redirect_uri} />
        </div>
        <div className="mb-4 md:flex md:items-center">
          <Label htmlFor="scope" />
          <Input name="scope" value={parameter.scope} />
        </div>
        <div className="mb-4 md:flex md:items-center">
          <Label htmlFor="state" />
          <Input name="state" value={parameter.state} />
        </div>
        <div className="mb-4 md:flex md:items-center">
          <Label htmlFor="response_type" />
          <Input name="response_type" value={parameter.response_type} />
        </div>
        <div className="mb-4 md:flex md:items-center">
          <Label htmlFor="consent" />
          <Input name="consent" value="true" />
        </div>
        <div className="mb-4 md:flex md:items-center">
          <input
            type="submit"
            className="rounded bg-purple-500 py-2 px-4 font-bold text-white shadow hover:bg-purple-400 focus:outline-none"
            value="Consent"
          />
        </div>
      </form>
    </div>
  )
}

const Input = (props: { name: string; value: string | number | undefined }) => {
  return (
    <input
      name={props.name}
      value={props.value}
      className="w-full appearance-none rounded border-2 border-gray-200 bg-gray-200 py-2 px-4 leading-tight text-gray-700 focus:border-purple-500 focus:bg-white focus:outline-none"
    />
  )
}

const Label = (props: { htmlFor: string }) => {
  return (
    <label htmlFor={props.htmlFor} className="mb-1 block pr-4 font-bold text-gray-500 md:mb-0">
      {`${props.htmlFor}: `}
    </label>
  )
}
export default AuthorizePage
