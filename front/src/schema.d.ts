/**
 * This file was auto-generated by openapi-typescript.
 * Do not make direct changes to the file.
 */

export interface paths {
  '/auth/login/user': {
    /** @description 顧客のログイン用エンドポイント(実装した) */
    post: operations['post-login'];
  };
  '/auth/signup/user': {
    /** @description 顧客のメールアドレスのリンクのtokenを送り本登録する(実装した) */
    get: operations['get-auth-signup'];
    /** @description 顧客の新規登録(実装した) */
    post: operations['post-auth-signup'];
  };
  '/auth/login/shop': {
    /** @description お店のログイン用エンドポイント */
    post: operations['post-api-auth-login-shop'];
  };
  '/auth/signup/shop': {
    /**
     * Your GET endpoint
     * @description 店のメールアドレスのリンクのtokenを送り本登録する
     */
    get: operations['get-auth-signup-shop'];
    /** @description 店の新規登録 */
    post: operations['post-auth-signup-shop'];
  };
  '/auth/logout': {
    /** @description ログアウト用のエンドポイント */
    post: operations['post-api-auth-logout'];
  };
  '/auth/password/reset/link': {
    get: operations['get-auth-password-reset-link'];
    post: operations['post-password-reset'];
  };
  '/users': {
    /** @description 顧客様一覧。認証店舗に紐づいた全てのユーザを返却する */
    get: operations['get-users'];
  };
  '/bottles': {
    /**
     * Your GET endpoint
     * @description ユーザが自分のキープボトル一覧を取得するときに使う
     */
    get: operations['get-bottles'];
    /** @description ボトル申請 */
    post: operations['post-bottles'];
  };
  '/shops/bottles': {
    /**
     * Your GET endpoint
     * @description 店舗がキープボトル一覧を取得するときに使う
     */
    get: operations['get-shops-bottles'];
  };
  '/shops': {
    /** Your GET endpoint */
    get: operations['get-shops'];
    put: operations['put-shops'];
  };
  '/me': {
    /** Your GET endpoint */
    get: operations['get-mypage'];
    put: operations['put-me'];
  };
  '/bottles/{id}': {
    get: operations['get-bottles-id'];
    put: operations['put-bottles-id'];
    delete: operations['delete-bottles-id'];
    parameters: {
      path: {
        id: string;
      };
    };
  };
  '/users/{id}': {
    /** Your GET endpoint */
    get: operations['get-users-id'];
    parameters: {
      path: {
        id: string;
      };
    };
  };
}

export type webhooks = Record<string, never>;

export interface components {
  schemas: {
    /**
     * user
     * @description お客様
     */
    user: {
      first_name: string;
      first_name_huri: string;
      last_name: string;
      last_name_huri: string;
      screen_name: string;
      /**
       * Format: date
       * @example 2019-08-24
       */
      birthday: string;
      /** Format: email */
      email: string;
    };
    /** bottle */
    bottle: {
      /** @enum {string} */
      type: 'whisky' | 'shochu' | 'brandy';
      name: string;
    };
    /** shop */
    shop: {
      name: string;
      name_kana: string;
      email: string;
      phone: string;
      owner: {
        name: string;
        name_kana: string;
        phone: string;
      };
    };
  };
  responses: never;
  parameters: never;
  requestBodies: never;
  headers: never;
  pathItems: never;
}

export type $defs = Record<string, never>;

export type external = Record<string, never>;

export interface operations {
  /** @description 顧客のログイン用エンドポイント(実装した) */
  'post-login': {
    requestBody?: {
      content: {
        'application/json': {
          email: string;
          password: string;
        };
      };
    };
    responses: {
      /** @description OK */
      200: {
        content: never;
      };
    };
  };
  /** @description 顧客のメールアドレスのリンクのtokenを送り本登録する(実装した) */
  'get-auth-signup': {
    parameters: {
      query: {
        token: string;
      };
    };
    responses: {
      /** @description OK */
      200: {
        content: never;
      };
    };
  };
  /** @description 顧客の新規登録(実装した) */
  'post-auth-signup': {
    requestBody?: {
      content: {
        'application/json': {
          ''?: components['schemas']['user'];
        } & {
          password: string;
          password_confirm: string;
        };
      };
    };
    responses: {
      /** @description OK */
      200: {
        content: never;
      };
    };
  };
  /** @description お店のログイン用エンドポイント */
  'post-api-auth-login-shop': {
    requestBody?: {
      content: {
        'application/json': {
          email?: string;
          password?: string;
        };
      };
    };
    responses: {
      /** @description OK */
      200: {
        content: never;
      };
    };
  };
  /**
   * Your GET endpoint
   * @description 店のメールアドレスのリンクのtokenを送り本登録する
   */
  'get-auth-signup-shop': {
    parameters: {
      query: {
        token: string;
      };
    };
    responses: {
      /** @description OK */
      200: {
        content: never;
      };
    };
  };
  /** @description 店の新規登録 */
  'post-auth-signup-shop': {
    requestBody?: {
      content: {
        'application/json': components['schemas']['shop'] & {
          password?: string;
          password_confirm?: string;
          opening_hours?: {
            ''?: string;
          };
        };
      };
    };
    responses: {
      /** @description OK */
      200: {
        content: never;
      };
    };
  };
  /** @description ログアウト用のエンドポイント */
  'post-api-auth-logout': {
    parameters: {
      cookie: {
        session_id: string;
      };
    };
    responses: {
      /** @description OK */
      200: {
        content: never;
      };
    };
  };
  'get-auth-password-reset-link': {
    parameters: {
      query: {
        token: string;
      };
    };
    responses: {
      /** @description OK */
      200: {
        content: never;
      };
    };
  };
  'post-password-reset': {
    requestBody?: {
      content: {
        'application/json': {
          email: string;
        };
      };
    };
    responses: {
      /** @description OK */
      200: {
        content: never;
      };
    };
  };
  /** @description 顧客様一覧。認証店舗に紐づいた全てのユーザを返却する */
  'get-users': {
    responses: {
      /** @description OK */
      200: {
        content: {
          'application/json': {
            users: {
              id: string;
              name: string;
              img: string;
              /**
               * Format: date
               * @example 2019-12-31
               */
              birthday: string;
            }[];
          };
        };
      };
    };
  };
  /**
   * Your GET endpoint
   * @description ユーザが自分のキープボトル一覧を取得するときに使う
   */
  'get-bottles': {
    parameters: {
      cookie?: {
        undefined?: string;
      };
    };
    responses: {
      /** @description OK */
      200: {
        content: {
          'application/json': {
            bottles: (components['schemas']['bottle'] & {
              /** Format: date */
              expires_at: string;
              /** Format: float */
              amount: number;
              status: {
                /** @enum {string} */
                status: 'approved' | 'rejected' | 'pending';
                /** @description 差し戻し理由 */
                reason?: string;
              };
              shop: {
                name: string;
                id: string;
              };
              id: string;
            })[];
          };
        };
      };
    };
  };
  /** @description ボトル申請 */
  'post-bottles': {
    requestBody?: {
      content: {
        'application/json': {
          bottles: (components['schemas']['bottle'] & {
            shop_id: string;
            /** Format: date */
            opened_at: string;
          })[];
        };
        'application/xml': {
          ''?: string;
        };
      };
    };
    responses: {
      /** @description OK */
      200: {
        content: never;
      };
    };
  };
  /**
   * Your GET endpoint
   * @description 店舗がキープボトル一覧を取得するときに使う
   */
  'get-shops-bottles': {
    responses: {
      /** @description OK */
      200: {
        content: {
          'application/json': {
            bottles: (components['schemas']['bottle'] & {
              /** Format: date */
              expires_at: string;
              /** Format: float */
              amount: number;
              id: string;
              username: string;
              /** @enum {string} */
              status: 'pending' | 'approved';
            })[];
          };
        };
      };
    };
  };
  /** Your GET endpoint */
  'get-shops': {
    responses: {
      /** @description OK */
      200: {
        content: {
          'application/json': {
            shops: {
              lat: number;
              lng: number;
              name: string;
              id: string;
              address: string;
            }[];
          };
        };
      };
    };
  };
  'put-shops': {
    requestBody?: {
      content: {
        'application/json': components['schemas']['shop'];
      };
    };
    responses: {
      /** @description OK */
      200: {
        content: never;
      };
    };
  };
  /** Your GET endpoint */
  'get-mypage': {
    requestBody?: {
      content: {
        'application/json': Record<string, never>;
      };
    };
    responses: {
      /** @description OK */
      200: {
        content: {
          'application/json': components['schemas']['user'] & {
            /** Format: uri */
            img?: string;
          };
        };
      };
    };
  };
  'put-me': {
    requestBody?: {
      content: {
        'application/json': {
          /** @description dataURL */
          img: string;
        };
      };
    };
    responses: {
      /** @description OK */
      200: {
        content: never;
      };
    };
  };
  'get-bottles-id': {
    parameters: {
      path: {
        id: string;
      };
    };
    responses: {
      /** @description OK */
      200: {
        content: {
          'application/json': components['schemas']['bottle'] & {
            /** Format: date */
            updated_at: string;
            /** @enum {unknown} */
            updated_by: 'shop' | 'user';
            /** Format: float */
            amount: number;
            shop: {
              name: string;
              address: string;
              phone: string;
              'eigyouzikan???': string;
            };
          };
        };
      };
    };
  };
  'put-bottles-id': {
    parameters: {
      path: {
        id: string;
      };
    };
    requestBody?: {
      content: {
        'application/json': {
          /** Format: float */
          amount?: number;
          /** @enum {string} */
          status?: 'approved' | 'rejected';
          reason?: string;
        };
      };
    };
    responses: {
      /** @description OK */
      200: {
        content: never;
      };
    };
  };
  'delete-bottles-id': {
    parameters: {
      path: {
        id: string;
      };
    };
    responses: {
      /** @description OK */
      200: {
        content: never;
      };
    };
  };
  /** Your GET endpoint */
  'get-users-id': {
    parameters: {
      path: {
        id: string;
      };
    };
    responses: {
      /** @description OK */
      200: {
        content: {
          'application/json': {
            name: string;
            img: string;
            /** Format: date */
            birthdate: string;
            bottles: (components['schemas']['bottle'] & {
              shop_name: string;
              /** Format: float */
              amount: number;
              /** Format: date */
              expires_at: string;
            })[];
          };
        };
      };
    };
  };
}
