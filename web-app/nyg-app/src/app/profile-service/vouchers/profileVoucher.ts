import { BehaviorSubject } from "rxjs"
import { v4 } from "uuid"

export interface PostProfilePattern {
    image: string
    comments: string
    alternative: string
}
export interface GetProfilePattern {
    imageUrl: string
    profileName: string
    nickname: string
}
export class ProfileFields {

    protected receivedTokens = new BehaviorSubject<GetProfilePattern>({ imageUrl: "", profileName: "", nickname: "" })

    protected Conn: EventSource

    protected URLs: string[] = ['http://localhost:5556/img/update']
    protected URLp: string = "http://localhost:8000/home/post/profile"
    protected URLg: string = "http://localhost:8000/home/get/profile"

}