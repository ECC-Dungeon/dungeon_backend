import{z as o,a as i}from"./index-ktjJghY_.js";import{u as s,a as u}from"./api-client-CipG6Tld.js";o.object({clear_floor:o.number().min(1,"Required")});const l=async({data:e})=>await u.post("/game/next",{clear_floor:e.clear_floor},{headers:{"Content-Type":"application/json",Authorization:localStorage.getItem("token")||""}}),p=({mutationConfig:e}={})=>{const r=i(),{onSuccess:t,...a}=e||{};return s({onSuccess:(...n)=>{r.invalidateQueries({queryKey:["next"]}),t==null||t(...n)},...a,mutationFn:l})};export{p as u};