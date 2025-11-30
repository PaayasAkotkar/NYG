
export interface IProfile {
    name: string
    nickname: string
    record: string
    gamesPlayed: number
    tier: string
    rating: number
    points: number
}
export interface IDeck {
    id: string
    _first: string
    _second: string
    isDefault: boolean
}
export interface IPowerUp {
    id: number
    covert: number
    nexus: number
    freeze: number
    rewind: number
    draw: number
    tag: number
    bet: number
}
export interface IIMG {
    id: string
    image: FormData
}


export interface IName {
    id: string
    name: string
}
export interface INickname {
    id: string
    nickname: string
}
