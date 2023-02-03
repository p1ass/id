'use client'
import { useSearchParams } from 'next/navigation'

import { AuthorizeRequestSchema } from '../../../../lib/oauth2/types'
import { buildAuthorizePath } from '../../../../pages/api/oauth2/route'

const labelClassName = 'mb-1 block pr-4 font-bold text-gray-500 md:mb-0'
const inputClassName =
  'w-full appearance-none rounded border-2 border-gray-200 bg-gray-200 py-2 px-4 leading-tight text-gray-700 focus:border-purple-500 focus:bg-white focus:outline-none'

const AuthorizePage = () => {
  const searchParams = useSearchParams()

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
          <label className={labelClassName} htmlFor="client_id">
            client_id:{' '}
          </label>
          <input name="client_id" value={parameter.client_id} className={inputClassName} />
        </div>
        <div className="mb-4 md:flex md:items-center">
          <label className={labelClassName} htmlFor="redirect_uri">
            redirect_uri:{' '}
          </label>
          <input name="redirect_uri" value={parameter.redirect_uri} className={inputClassName} />
        </div>
        <div className="mb-4 md:flex md:items-center">
          <label htmlFor="scope" className={labelClassName}>
            scope:{' '}
          </label>
          <input name="scope" value={parameter.scope} className={inputClassName} />
        </div>
        <div className="mb-4 md:flex md:items-center">
          <label htmlFor="state" className={labelClassName}>
            state:{' '}
          </label>
          <input name="state" value={parameter.state} className={inputClassName} />
        </div>
        <div className="mb-4 md:flex md:items-center">
          <label htmlFor="response_type" className={labelClassName}>
            response_type:{' '}
          </label>
          <input name="response_type" value={parameter.response_type} className={inputClassName} />
        </div>
        <div className="mb-4 md:flex md:items-center">
          <label htmlFor="consent" className={labelClassName}>
            consent:{' '}
          </label>
          <input name="consent" value="true" className={inputClassName} />
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

export default AuthorizePage
