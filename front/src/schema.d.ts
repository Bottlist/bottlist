/**
 * This file was auto-generated by openapi-typescript.
 * Do not make direct changes to the file.
 */

export interface paths {
  '/login': {
    post: operations['post-login'];
  };
  '/password/reset/link': {
    post: operations['post-password-reset'];
  };
  '/users': {
    get: operations['get-users'];
    /** @description 新規お客様登録 */
    post: operations['post-users'];
  };
  '/bottles': {
    /** Your GET endpoint */
    get: operations['get-bottles'];
    /** @description ボトル申請 */
    post: operations['post-bottles'];
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
      family_name: string;
      family_name_kana: string;
      last_name: string;
      last_name_kana: string;
      nickname: string;
      /** Format: date */
      birthdate: string;
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
            }[];
          };
        };
      };
    };
  };
  /** @description 新規お客様登録 */
  'post-users': {
    requestBody?: {
      content: {
        'application/json': components['schemas']['user'] & {
          password: string;
          password_confirm: string;
        };
        'application/xml': Record<string, never>;
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
  'get-bottles': {
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
  /** Your GET endpoint */
  'get-shops': {
    responses: {
      /** @description OK */
      200: {
        content: {
          'application/json': {
            shops: {
              lat: string;
              lng: string;
              name: string;
              id: string;
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
        'application/json': components['schemas']['bottle'] & {
          /** Format: float */
          amount?: number;
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