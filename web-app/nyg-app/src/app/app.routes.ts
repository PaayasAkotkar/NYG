import { Routes } from '@angular/router';
import { AppComponent } from './app.component';
import { LobbyComponent } from './lobby/lobby.component';
import { PartyRoomComponent } from './party-room/party-room.component';
import { LobbyUiComponent } from './UI/lobby-ui/lobby-ui.component';
import { JoinRoomComponent } from './UI/join-room/join-room.component';
import { CreateRoomComponent } from './UI/create-room/create-room.component';
import { SetupComponent } from './setup/setup.component';
import { MakeRoomComponent } from './make-room/make-room.component';
import { RoomsListComponent } from './UI/rooms-list/rooms-list.component';
import { HomePageComponent } from './home-page/home-page.component';
import { ClashLobbyComponent } from './UI/clash-lobby/clash-lobby.component';
import { LocifyLobbyComponent } from './UI/locify-lobby/locify-lobby.component';
import { LoginPageComponent } from './UI/login-page/login-page.component';
import { DeckLayoutComponent } from './UI/deck-layout/deck-layout.component';
import { Observable } from 'rxjs';

export const routes: Routes = [
    { path: 'Make', component: MakeRoomComponent, },
    { path: 'Join', component: LocifyLobbyComponent },
    { path: 'Home', component: HomePageComponent },
    { path: 'Game-Hall', component: LobbyComponent },
    { path: 'Search', component: JoinRoomComponent },
    { path: '', component: HomePageComponent },
    { path: 'Party-Hall', component: PartyRoomComponent },
    { path: 'Settings/:page', component: SetupComponent },
    { path: 'ClashLobby', component: LobbyComponent },
    { path: 'Clash', component: DeckLayoutComponent },
    { path: 'SignUp', component: LoginPageComponent },
    { path: 'Locify', component: LocifyLobbyComponent },
    { path: '___', component: LoginPageComponent },
    { path: '**', redirectTo: '' },
]

export interface MutextLock {
    deactivateGuard: () => Observable<boolean> | boolean;
}