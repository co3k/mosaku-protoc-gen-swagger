/// <reference path="types.ts"/>
/** @module LibraryService */
// Auto-generated, edits will be overwritten
import * as gateway from './gateway'

/**
 * @param {string} book.name 
 * @param {module:types.v1Book} body 
 * @return {Promise<module:types.v1Book>} 
 */
export function UpdateBook(book.name: string, body: api.v1Book): Promise<api.Response<api.v1Book>> {
  const parameters: api.OperationParamGroups = {
    path: {
      'book.name': book.name
    },
    body: {
      body
    }
  }
  return gateway.request(UpdateBookOperation, parameters)
}

/**
 * @param {string} name 
 * @return {Promise<module:types.v1Book>} 
 */
export function GetBook(name: string): Promise<api.Response<api.v1Book>> {
  const parameters: api.OperationParamGroups = {
    path: {
      name
    }
  }
  return gateway.request(GetBookOperation, parameters)
}

/**
 * @param {string} name 
 * @return {Promise<module:types.protobufEmpty>} 
 */
export function DeleteBook(name: string): Promise<api.Response<api.protobufEmpty>> {
  const parameters: api.OperationParamGroups = {
    path: {
      name
    }
  }
  return gateway.request(DeleteBookOperation, parameters)
}

/**
 * ListBooks は本を列挙します
 * 
 * @param {string} parent 
 * @param {object} options Optional options
 * @param {number} [options.pageSize] The maximum number of items to return.
 * @param {string} [options.pageToken] The next_page_token value returned from a previous List request, if any.
 * @return {Promise<module:types.v1ListBooksResponse>} 
 */
export function ListBooks(parent: string, options?: ListBooksOptions): Promise<api.Response<api.v1ListBooksResponse>> {
  if (!options) options = {}
  const parameters: api.OperationParamGroups = {
    path: {
      parent
    },
    query: {
      page_size: options.pageSize,
      page_token: options.pageToken
    }
  }
  return gateway.request(ListBooksOperation, parameters)
}

/**
 * @param {string} parent 
 * @param {module:types.v1Book} body 
 * @return {Promise<module:types.v1Book>} 
 */
export function CreateBook(parent: string, body: api.v1Book): Promise<api.Response<api.v1Book>> {
  const parameters: api.OperationParamGroups = {
    path: {
      parent
    },
    body: {
      body
    }
  }
  return gateway.request(CreateBookOperation, parameters)
}

export interface ListBooksOptions {
  /**
   * The maximum number of items to return.
   */
  pageSize?: number
  /**
   * The next_page_token value returned from a previous List request, if any.
   */
  pageToken?: string
}

const UpdateBookOperation: api.OperationInfo = {
  path: '/v1/{book.name}',
  contentTypes: ['application/json'],
  method: 'patch',
  security: [
    {
      id: 'OAuth2'
    }
  ]
}

const GetBookOperation: api.OperationInfo = {
  path: '/v1/{name}',
  method: 'get',
  security: [
    {
      id: 'OAuth2'
    }
  ]
}

const DeleteBookOperation: api.OperationInfo = {
  path: '/v1/{name}',
  method: 'delete',
  security: [
    {
      id: 'OAuth2'
    }
  ]
}

const ListBooksOperation: api.OperationInfo = {
  path: '/v1/{parent}/books',
  method: 'get',
  security: [
    {
      id: 'OAuth2'
    }
  ]
}

const CreateBookOperation: api.OperationInfo = {
  path: '/v1/{parent}/books',
  contentTypes: ['application/json'],
  method: 'post',
  security: [
    {
      id: 'OAuth2'
    }
  ]
}
