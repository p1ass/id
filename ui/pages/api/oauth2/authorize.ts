import type { NextApiRequest, NextApiResponse } from 'next'

export default function handler(req: NextApiRequest, res: NextApiResponse) {
  if (req.method === 'GET') {
    // TODO: implement. Copy from app/oauth2/authorize
    // If user consent is needed, redirect to consent page.
    // Otherwise, redirect to redirectUri
    return res.status(302).redirect('')
  }
  if (req.method === 'POST') {
    // TODO: implement
    return res.status(302).redirect('')
  }

  return res.status(405).json('{"message": "method not allowed"}')
}
