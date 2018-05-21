/// <reference path="../types.ts"/>
// Auto-generated, edits will be overwritten
const spec: api.OpenApiSpec = {
  'host': 'localhost',
  'schemes': [
    'https'
  ],
  'basePath': '',
  'contentTypes': [
    'application/json'
  ],
  'accepts': [
    'application/json'
  ],
  'securityDefinitions': {
    'OAuth2': {
      'type': 'oauth2',
      'flow': 'accessCode',
      'authorizationUrl': 'https://example.com/oauth/authorize',
      'tokenUrl': 'https://example.com/oauth/token'
    }
  }
}
export default spec
