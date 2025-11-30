import { FormBuilder, FormGroup } from "@angular/forms"
import { OCSFieldFormat } from "./patterns"
export interface GoBooleanMap {
    [key: string]: boolean
}

export namespace ColorPalletes {
    export const $nyg_green_palletes = {
        moss_green: '#262706',
        moss_green_light_shade: '#7F762B',
        moss_green_dark_shade: '#292207',
        olive_green: '#78A14D',
        olive_green_dark_shade: '#789D52',
        cobalt_green_light_shade: '#034831',
        cobalt_green: '#253A36',
        cobalt_green_feint_shade: '#DBE4E1',
        permanent_sap_green: '#292931',
        permanent_sap_green_light_shade: '#143E31',
        cadium_green: '#719502',
        cadium_green_light_shade: '#869F1F',
        italian_green_umber: '#66622E',
        italian_green_umber_feint_shade: '#808045',
        italian_green_umber_dark_shade: '#534927',
        potrait1_green: '#987011',
        potrait2_green: '#445529',
        potrait3_green: '#443321',
        potrait4_green: '#1F160D',
        potrait5_green: '#00E606',
        potrait6_green: '#0A3B00',
        potrait7_green: '#00D16B',
        ad1_green: '#71D702',
        ad2_green: '#7CAF3D',
        ad3_green: '#0A2F08',
        ad4_green: '#204009',
        ad5_green: '#93E800',
        ad6_green: '#EFFB00',
        ad7_green: '#CFFB49',
        ad8_green: '#006B33',
        ad9_green: '#5F6200',
        ad10_green: '#CFFB25',
        ad11_green: '#00110E',
        ad12_green: '#00FF43',
        ad13_green: '#BBC30E',
        ad14_green: '#96CD21',
        ad15_green: '#008D28',
        ad16_green: '#92FF1E',
        ad17_green: '#005E22',
        ad18_green: '#79FD00',
        ad19_green: '#74DB00',
        ad20_green: '#002B26',
    }

    export const $nyg_white_palletes = {
        warm_white_alternative: '#F1EFE1',
        warm_white_light_shade: '#FAF8EA',
        white: '#EFEFEF',
        creamize_white: '#FAF7FC',
        foundation_white: '#F2EAE8',
        warm_white: '#C3B7AF',
        tatanium_white: '#F0ECF0',
        ad1_white: '#E5E3E1',
        ad2_white: '#45556B',
        ad3_white: '#5F7493',
        ad4_white: '#EFFDBC',
        ad5_white: '#C3B9B9',
        ad6_white: '#AFAFAF',
        ad7_white: '#CED0D0',
        ad8_white: '#F5F4DE',
        ad9_white: '#262626',
    }

    export const $nyg_red_palletes = {
        scarllet: '#A1011C',
        dore_rose: '#8D2733',
        dore_rose_light_shade: '#C7A5A6',
        dore_rose_slight_light_shade: '#90182B',
        indian_red: '#BD3851',
        alzarian_caret: '#80000D',
        potrait1_red: '#501916',
        potrait_red_light_shade: '#501916',
        potrait_red_medium_shade: '#631E1B',
        potrait2_red: '#732C1E',
        potrait3_red: '#5A1F1B',
        potrait4_red: '#6C2B02',
        potrait5_red: '#621512',
        potrait6_red: '#B47B6C',
        potrait7_red: '#88232A',
        potrait8_red: '#6D1114',
        potrait9_red: '#DC3233',
        potrait10_red: '#730000',
        ad1_red: '#D2011C',
        ad2_red: '#B60120',
        ad3_red: '#FF47FF',
        ad4_red: '#E64265',
        ad5_red: '#DC518C',
        ad6_red: '#870260',
        ad7_red: '#740707',
        ad8_red: '#B60000',
        ad9_red: '#8D003B',
        ad10_red: '#960000',
        ad11_red: '#FF0060',
        ad12_red: '#E80014',
        ad13_red: '#BD0000',
        ad14_red: '#FF3497',
        ad15_red: '#FF0044',
        ad16_red: '#FF0034',
        ad17_red: '#B10000'
    }
    export const $nyg_blue_palletes = {
        cerulean: '#0584C6',
        cerulean_dark_shade: '#0F5891',
        kings_blue: '#5395E9',
        vivid_blue: '#66AABA',
        potrait1_blue: '#1A2C32',
        potrait2_blue: '#D6DCE8',
        potrait3_blue: '#0F607A',
        potrait4_blue: '#1B292A',
        potrait5_blue: '#025BE0',
        potrait6_blue: '#2D0DD3',
        potrait7_blue: '#0068FF',
        potrait8_blue: '#227C71',
        potrait9_blue: '#2D294C',
        ad1_blue: '#22A971',
        ad2_blue: '#1A2CC2',
        ad3_blue: '#1vad6C2',
        ad4_blue: '#8181FF',
        ad5_blue: '#03ACAC',
        ad6_blue: '#0016A4',
        ad7_blue: '#00FFE2',
        ad8_blue: '#009ACE',
        ad9_blue: '#007DB4',
        ad10_blue: '#0091FF',
        ad11_blue: '#6B05FF',
        ad12_blue: '#7434FF',
        ad13_blue: '#004CFF',
        ad14_blue: '#008777',
        ad15_blue: '#26264A',
        ad16_blue: '#26268B',
        ad17_blue: '#003C30',
    }
    export const $nyg_black_palletes = {
        black: '#1B1714',
        umber_black: '#252A19',
        potrait1_black: '#1D2225',
        potrait2_black: '#37302E',
        potrait3_black: '#010004',
        potrait4_black: '#241713',
        ad1_black: '#01091B',
        ad2_black: '#0E091B',
        ad3_black: '#140210',
        ad4_black: '#252112',
        ad5_black: '#0E0F0D',
        ad6_black: '#000922',
        ad7_black: '#323032',
        ad8_black: '#0E0F0D',
    }
    export const $nyg_yellow_palletes = {
        yellow_orch: '#AE9A46',
        turner_yellow: '#E5B62B',
        turner_yellow_dark_shvade: '#DB9412',
        turner_yellow_feint_shvade: '#E6DBB5',
        potrait1_yellow_feint_shade: '#D0B47F',
        potrait2_yellow: '#CDA23D',
        potrait3_yellow: '#795F3E',
        potrait4_yellow: '#F0B81D',
        potrait5_yellow: '#EACA28',
        potrait6_yellow: '#F0D500',
        ad1_yellow: '#FFB02D',
        ad2_yellow: '#FFF756',
        ad3_yellow: '#C2BD00',
        ad4_yellow: '#FFC30E',
        ad5_yellow: '#FFF086',
        ad6_yellow: '#ECF67D',
        ad7_yellow: '#EEFF34',
        ad8_yellow: '#FFFF00',
        ad9_yellow: '#FFFB7E',
        ad10_yellow: '#FFFF55'
    }

    export const $nyg_brown_palletes = {
        brown: '#6A3F25',
        brunt_umber: '#3D2322',
        brunt_umber_bright_shade: '#6F250E',
        potrait2_brown: '#834E22',
        ad1_brown: '#6C2C10',
        ad2_brown: '#44170A',
        ad3_brown: '#815700',
        ad4_brown: '#790000',
        ad5_brown: '#1F0000',
        ad7_brown: '#200C00',
        ad8_brown: '#211400'
    }
    export const $nyg_voilet_palletes = {
        magenta_light_shade: '#A85387',
        magenta_dark_shade: '#622042',
        magenta_feint_shade: '#BBABBA',
        magenta_dark_shade_2: '#691425',
        magenta_medium_shade: '#792B3A',
        potrait1_voilet: '#4D4051',
        ad1_voilet: '#3E0C76',
        ad2_voilet: '#2D0A27',
        potrait2_voilet: '#2E2144',
        ad3_voilet: '#A33DF8',
        ad4_voilet: '#FF1425',
        ad5_voilet: '#6F09E8',
        ad6_voilet: '#54003E',
        ad7_voilet: '#FF0012',
        ad8_voilet: '#750062',
        ad9_voilet: '#FF91FF',

    }

    export const $nyg_orange_palletes = {
        cadium_orange: '#F29136',
        flame_orange: '#CD7237',
        // or FFB02D
        flame_orange_dark_shade: '#C8632B',
        flame_orange_bright_shade: ', #BF4822',
        skin_tone: '#C0824C',
        skin_tone2: '#B66836',
        skin_tone3: '#B25337',
        skin_tone4: '#A05423',
        skin_tone5: '#9B7A41',
        skin_tone6: '#DFD1AB',
        potrait1_orange: '#D64B00',
        potrait2_orange: '#B57324',
        ad1_orange: '#FFA800',
        ad2_orange: '#D69E00',
        ad3_orange: '#FF2B29',
        ad4_orange: '#FF981F',
        ad5_orange: '#FFC000',
        ad6_orange: '#FF6F00',
        ad7_orange: '#FF9000',
        ad8_orange: '#FF0700',
        ad9_orange: '#FF3C23',
    }
}
export namespace HAND {
    /**
     * 
     * @param threeD three D array string
     * @returns converts the 3d arr to 2d arr
     */
    export function ConvTwoD(threeD: string[][][]): string[][] {
        var twoD: string[][] = []
        for (let item of threeD) {
            for (let itemx of item) {
                twoD.push(itemx)
            }
        }
        return twoD
    }

    export function ConvDtoTwoD(oneD: string[]): string[][] {
        var twoD = []
        for (let i of oneD) {
            twoD.push([i])
        }
        return twoD
    }
    /**
     * 
     * @param twoD two D array string
     * @returns converts the 2d arr to 3d arr
     */
    export function ConvThreeD(twoD: string[][]): string[][][] {
        var threeD = [twoD]

        return threeD
    }

    /**
    * 
    * @param  sentencesSep 2d array sentences 
    * @returns sepration of words in the sentences
    * @example input: [['king is lazy']] output: [['king','is','lazy']]
    */
    export function sentencesSeparation(sentencesSep: string[][], removeNum: boolean): string[][] {
        var IA, sentences, res = []
        for (let i = 0; i < sentencesSep.length; i++) {
            IA = sentencesSep[i]
            sentences = IA[0] // imp 
            res.push(sentences.split(' '))
        }
        var temp: any = []
        if (removeNum) {
            res.forEach((e) => {
                temp.push(e.filter((x) => { return /[a-zA-z]/.test(x) }))
            })
            res = []
            res = temp
        }
        return res
    }
    /**
     * @param words list of sentences 
     * @param ws to remove whitespace
     * @return collection of separated letters
     * @example input:['apple', 'mango'] output: [['a','p','p','l','e'],[....]] note: if the ws is true
     */


    export function wordSeparation(words: string[], ws: boolean = false): string[][] {
        var _______C: string[][] = []

        if (ws) {
            for (let i of words) {
                _______C.push(i.split('').filter(e => e != ' '))
            }
        } else {
            for (let i of words) {
                _______C.push(i.split(''))
            }

        }
        return _______C
    }

    export class Queue {
        store: string[] = []
        len: number = 0
        /**
         * 
         * @param d stores the element
         */
        enqueue(d: string) {
            this.store.push(d)
            this.len = this.store.length
        }

        Range(): string[] {
            return this.store
        }
        /**
         * 
         * @returns the front element
         */
        front() {

            if (this.len == 0) {
                this.len = this.store.length
            } else {
                this.len += this.store.length
            }

            return this.store[0]
        }
        /**
         * removes the newly added element
         */
        dequeue() {
            return this.store.shift()
        }
        isEmpty() {
            return this.store.length == 0
        }
        constructor() { }
    }

    /**
     * 
     * @param words array of sentences 
     * @returns separates the sentences of words into a double array
     * @example input: ['apple', 'mango'] ouput: [['apple'], ['mango']]
     */
    export function wordSeparationV2(words: string[], ws: boolean = false): string[][] {
        var ___C = words.map(e => e.split(' '))
        return ___C
    }

    /**
     * 
     * @param sentence data of names ; for example ['a', 'a b', 'ab c']
     * @param ws remove whitespace
     * @param single returns only non-whitespace character
     * @param double returns only whitespace characters
     * note: if you set single and double false it would return [single, double]
     * @returns converts the non-whitespace to all the string while converts the whitespace sentence to group string
     * @example input: ['apple is healty', 'ab', 'catch us'] output: [['apple', 'is', 'healty'], ['a', 'b'], ['catch', 'us']] note: if ws 
     * 
     * 
     */
    // input :[a , a c , ab ] output: [[a], [a, c], [a, b]]
    export function WhiteSentenceSeparation(sentence: string[], ws: boolean = false, single: boolean = false, double: boolean = false): string[][] {
        var __FOUNND: string[] = sentence.filter(e => /\s/.test(e))
        var __NOT: string[] = sentence.filter(e => !/\s/.test(e))
        var F: string[][], N: string[][]
        F = sentencesSeparation(ConvDtoTwoD(__FOUNND), ws)
        F = wordSeparationV2(__FOUNND)
        N = wordSeparation(__NOT, ws)

        var __done
        if (single) {
            __done = N
        } else if (double) {
            __done = F
        } else {
            __done = [F, N]
            __done = ConvTwoD(__done)
        }
        return __done
    }

    export function FillControls(FillForm: FormGroup, fillForm: OCSFieldFormat[], FB: FormBuilder): void {
        fillForm.forEach(item => {
            FillForm.addControl<any>(item.OCStemplate, FB.control(item.OCSformControl))
        })
    }

    // Calculate adds the given number with string 
    // and than appends with the matched regex token
    export function Calculate(str: string, add: number, re: RegExp): string {
        var temp: string = ""
        for (let i = 0; i < str.length; i++) {
            if (re.test(str[i])) {
                temp = str[i]
                str = str.slice(0, i)
                break
            }
        }
        str = String(Number.parseInt(str) + add) + temp
        return str
    }

}


