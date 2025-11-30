import { ApplicationConfig, provideZoneChangeDetection, inject, Inject } from '@angular/core';
import { provideRouter } from '@angular/router';

import { routes } from './app.routes';
import { provideClientHydration, withEventReplay } from '@angular/platform-browser';
import { HttpClient, provideHttpClient, withFetch } from '@angular/common/http';
import { provideServerRendering } from '@angular/platform-server';
import { provideAnimations, provideNoopAnimations } from '@angular/platform-browser/animations';
import { provideApollo, provideNamedApollo } from 'apollo-angular';
import { HttpLink } from 'apollo-angular/http';
import { InMemoryCache, split } from '@apollo/client/core';
import { GraphQLWsLink } from "@apollo/client/link/subscriptions";
import { createClient } from "graphql-ws";
import { getMainDefinition } from '@apollo/client/utilities';
import { Kind, OperationTypeNode } from 'graphql';

const init = () => {
  const http = inject(HttpLink)
  const hclientSettingsChanges = 'http://localhost:6060/settings-changes'
  const hclientProfile = 'http://localhost:6060/profile'
  const hclientBooks = 'http://localhost:6060/books'
  const hclientLogin = 'http://localhost:6060/login'
  const hclientChatInit = 'http://localhost:1299/nyg-chat-init'
  const hclientChatLatestMsg = 'http://localhost:1299/nyg-chat-latest-msg'


  const hcsc = http.create({ uri: hclientSettingsChanges, })
  const hcp = http.create({ uri: hclientProfile })
  const hcb = http.create({ uri: hclientBooks })
  const hcl = http.create({ uri: hclientLogin })
  const hcci = http.create({ uri: hclientChatInit })
  const hclm = http.create({ uri: hclientChatLatestMsg })

  const wsclientSettingsChanges = createClient({ url: 'ws://localhost:6060/settings-changes', connectionParams: { authToken: 'NYG' } },)
  const wsclientProfile = createClient({ url: 'ws://localhost:6060/profile' })
  const wsclientBooks = createClient({ url: 'ws://localhost:6060/books' })
  const wsclientLogin = createClient({ url: 'ws://localhost:6060/login' })
  const wsclientLatestInit = createClient({ url: 'ws://localhost:1299/nyg-chat-init' })
  const wsclientLatestMessage = createClient({ url: 'ws://localhost:1299/nyg-chat-latest-msg' })

  const wcsc = new GraphQLWsLink(wsclientSettingsChanges)
  const wcp = new GraphQLWsLink(wsclientProfile)
  const wcb = new GraphQLWsLink(wsclientBooks)
  const wcl = new GraphQLWsLink(wsclientLogin)
  const wclm = new GraphQLWsLink(wsclientLatestInit)
  const wci = new GraphQLWsLink(wsclientLatestMessage)

  const cscLink = split(
    ({ query }) => {
      const def = getMainDefinition(query)
      return def.kind === Kind.OPERATION_DEFINITION &&
        def.operation === OperationTypeNode.SUBSCRIPTION
    },
    wcsc,
    hcsc
  )
  const cpLink = split(
    ({ query }) => {
      const def = getMainDefinition(query)
      return def.kind === Kind.OPERATION_DEFINITION &&
        def.operation === OperationTypeNode.SUBSCRIPTION
    },
    wcp,
    hcp
  )
  const cbLink = split(
    ({ query }) => {
      const def = getMainDefinition(query)
      return def.kind === Kind.OPERATION_DEFINITION &&
        def.operation === OperationTypeNode.SUBSCRIPTION
    },
    wcb,
    hcb
  )
  const clLink = split(
    ({ query }) => {
      const def = getMainDefinition(query)
      return def.kind === Kind.OPERATION_DEFINITION &&
        def.operation === OperationTypeNode.SUBSCRIPTION
    },
    wcl,
    hcl
  )
  const ccciLink = split(
    ({ query }) => {
      const def = getMainDefinition(query)
      return def.kind === Kind.OPERATION_DEFINITION &&
        def.operation === OperationTypeNode.SUBSCRIPTION
    },
    wclm,
    hcci
  )
  const cclmLink = split(
    ({ query }) => {
      const def = getMainDefinition(query)
      return def.kind === Kind.OPERATION_DEFINITION &&
        def.operation === OperationTypeNode.SUBSCRIPTION
    },
    wci,
    hclm
  )

  return {
    setup_settings: {
      link: cscLink,
      cache: new InMemoryCache(),
    },
    profile: {
      link: cpLink,
      cache: new InMemoryCache(),
    },
    books: {
      link: cbLink,
      cache: new InMemoryCache(),
    },
    login: {
      link: clLink,
      cache: new InMemoryCache(),
    },
    chat_room: {
      link: cclmLink,
      cache: new InMemoryCache(),
    },
    init_room: {
      link: ccciLink,
      cache: new InMemoryCache(),
    }
  }

}

export const appConfig: ApplicationConfig = {
  providers: [provideZoneChangeDetection({ eventCoalescing: true }), provideRouter(routes), provideClientHydration(withEventReplay(),)
    , provideHttpClient(withFetch()), provideClientHydration(), provideAnimations(), provideHttpClient(),

  provideNamedApollo(init),
  ]
}



// const httpLink = inject(HttpLink);

// return {
//   setup_settings: {
//     link: httpLink.create({
//       uri: 'http://localhost:6060/settings-changes',
//     }),
//     cache: new InMemoryCache(),
//   },
//   profile: {
//     link: httpLink.create({
//       uri: 'http://localhost:6060/profile',
//     }),
//     cache: new InMemoryCache(),
//   },
//   books: {
//     link: httpLink.create({
//       uri: 'http://localhost:6060/books',
//     }),
//     cache: new InMemoryCache(),
//   },

// }
// () => {

//   const clientSettingsChanges = createClient({ url: 'ws://localhost:6060/settings-changes' })
//   const clientProfile = createClient({ url: 'ws://localhost:6060/profile' })
//   const clientBooks = createClient({ url: 'ws://localhost:6060/books' })
//   return {
//     setup_settings: {
//       link: new GraphQLWsLink(clientSettingsChanges),
//       cache: new InMemoryCache(),
//     },
//     profile: {
//       link: new GraphQLWsLink(clientProfile),
//       cache: new InMemoryCache(),
//     },
//     books: {
//       link: new GraphQLWsLink(clientBooks),
//       cache: new InMemoryCache(),
//     },

//   }
// }