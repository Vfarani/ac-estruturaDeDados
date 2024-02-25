package main

func CalculaMedia(nums ...float64) float64 {

    if len(nums) == 0 {
        return 0
    }

    soma := 0.0
    for _, num := range nums {
        soma += num
    }

    media := soma / float64(len(nums))
    return media
}
