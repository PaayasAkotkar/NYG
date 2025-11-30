import { ComponentFixture, TestBed } from '@angular/core/testing'
import { LobbyComponent } from './lobby.component'
import { Apollo } from 'apollo-angular'
import { ApolloTestingModule } from 'apollo-angular/testing'

describe('LobbyComponent', () => {
  let component: LobbyComponent
  let fixture: ComponentFixture<LobbyComponent>

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [LobbyComponent],
      providers: [Apollo, ApolloTestingModule]
    })
      .compileComponents()

    fixture = TestBed.createComponent(LobbyComponent)
    component = fixture.componentInstance
    fixture.detectChanges()
  })

  it('should create', () => {
    expect(component).toBeTruthy()
  })

})
